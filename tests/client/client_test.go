package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/grpclb"
	"github.com/eviltomorrow/robber-datasource/internal/server"
	"github.com/eviltomorrow/robber-datasource/pkg/pb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestVersion(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
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
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = cli.Status(ctx, "localhost:2379")
	if err != nil {
		log.Fatal(err)
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
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Println("connetion ...")

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := pb.NewServiceClient(conn)
	repley, err := client.Version(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Version error: %v", err)
	}
	fmt.Println(repley.Value)
}
