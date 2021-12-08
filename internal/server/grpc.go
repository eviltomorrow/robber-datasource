package server

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/grpclb"
	"github.com/eviltomorrow/robber-core/pkg/mongodb"
	"github.com/eviltomorrow/robber-core/pkg/system"
	"github.com/eviltomorrow/robber-core/pkg/zlog"
	"github.com/eviltomorrow/robber-core/pkg/znet"
	"github.com/eviltomorrow/robber-datasource/internal/middleware"
	"github.com/eviltomorrow/robber-datasource/internal/service"
	"github.com/eviltomorrow/robber-datasource/pkg/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	Host           = "0.0.0.0"
	Port           = 19090
	Endpoints      = []string{}
	RevokeEtcdConn func() error
	Key            = "grpclb/service/database"

	server *grpc.Server
)

type GRPC struct {
	pb.UnimplementedServiceServer
}

func (g *GRPC) Version(ctx context.Context, _ *emptypb.Empty) (*wrapperspb.StringValue, error) {
	var buf bytes.Buffer
	buf.WriteString("Server: \r\n")
	buf.WriteString(fmt.Sprintf("   Robber-database-sina Version (Current): %s\r\n", system.MainVersion))
	buf.WriteString(fmt.Sprintf("   Go Version: %v\r\n", system.GoVersion))
	buf.WriteString(fmt.Sprintf("   Go OS/Arch: %v\r\n", system.GoOSArch))
	buf.WriteString(fmt.Sprintf("   Git Sha: %v\r\n", system.GitSha))
	buf.WriteString(fmt.Sprintf("   Git Tag: %v\r\n", system.GitTag))
	buf.WriteString(fmt.Sprintf("   Git Branch: %v\r\n", system.GitBranch))
	buf.WriteString(fmt.Sprintf("   Build Time: %v\r\n", system.BuildTime))
	buf.WriteString(fmt.Sprintf("   HostName: %v\r\n", system.HostName))
	buf.WriteString(fmt.Sprintf("   IP: %v\r\n", system.IP))
	buf.WriteString(fmt.Sprintf("   Running Time: %v\r\n", system.RunningTime()))
	return &wrapperspb.StringValue{Value: buf.String()}, nil
}

func (g *GRPC) PullData(req *wrapperspb.StringValue, resp pb.Service_PullDataServer) error {
	var (
		offset  int64 = 0
		limit   int64 = 100
		lastID  string
		timeout = 20 * time.Second
	)

	for {
		data, err := service.SelectMetadataRange(mongodb.DB, offset, limit, req.Value, lastID, timeout)
		if err != nil {
			zlog.Error("SelectMetadataRange failure", zap.Error(err))
			break
		}
		for _, d := range data {
			err := resp.Send(&pb.Metadata{
				Code:            d.Code,
				Name:            d.Name,
				Open:            d.Open,
				YesterdayClosed: d.YesterdayClosed,
				Latest:          d.Latest,
				High:            d.High,
				Low:             d.Low,
				Volume:          d.Volume,
				Account:         d.Account,
				Date:            d.Date,
				Time:            d.Time,
				Suspend:         d.Suspend,
			})
			if err != nil {
				return err
			}
		}

		if len(data) < int(limit) {
			break
		}
		if len(data) != 0 {
			lastID = data[len(data)-1].ObjectID
		}
		offset += limit
	}
	return nil
}

func StartupGRPC() error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Host, Port))
	if err != nil {
		return err
	}

	server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.UnaryServerRecoveryInterceptor,
		),
	)

	reflection.Register(server)
	pb.RegisterServiceServer(server, &GRPC{})

	localIp, err := znet.GetLocalIP2()
	if err != nil {
		return fmt.Errorf("get local ip failure, nest error: %v", err)
	}

	close, err := grpclb.Register(Key, localIp, Port, Endpoints, 10)
	if err != nil {
		return fmt.Errorf("register service to etcd failure, nest error: %v", err)
	}
	RevokeEtcdConn = func() error {
		close()
		return nil
	}

	go func() {
		if err := server.Serve(listen); err != nil {
			zlog.Fatal("GRPC Server startup failure", zap.Error(err))
		}
	}()
	return nil
}

func ShutdownGRPC() error {
	if server == nil {
		return nil
	}
	server.Stop()
	return nil
}
