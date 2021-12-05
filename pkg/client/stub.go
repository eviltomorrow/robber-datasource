package client

import (
	"context"
	"fmt"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/grpclb"
	"github.com/eviltomorrow/robber-datasource/internal/server"
	"github.com/eviltomorrow/robber-datasource/pkg/pb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

var (
	EtcdEndpoints = []string{
		"localhost:2379",
	}
)

func NewClientForDatasource() (pb.ServiceClient, func(), error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   EtcdEndpoints,
		DialTimeout: 5 * time.Second,
		LogConfig: &zap.Config{
			Level:            zap.NewAtomicLevelAt(zap.ErrorLevel),
			Development:      false,
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	builder := &grpclb.Builder{
		Client: cli,
	}
	resolver.Register(builder)

	target := fmt.Sprintf("etcd:///%s", server.Key)
	conn, err := grpc.DialContext(
		context.Background(),
		target,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, nil, err
	}
	return pb.NewServiceClient(conn), func() { conn.Close() }, nil
}
