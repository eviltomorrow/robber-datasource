package middleware

import (
	"context"
	"path"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/zlog"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// UnaryServerLogInterceptor log 拦截
func UnaryServerLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	var addr string
	if peer, ok := peer.FromContext(ctx); ok {
		addr = peer.Addr.String()
	}

	var start = time.Now()
	defer func() {
		zlog.Info("grpc(*unary)",
			zap.String("service", path.Dir(info.FullMethod)[1:]),
			zap.String("method", path.Base(info.FullMethod)),
			zap.String("addr", addr),
			zap.Any("req", req),
			zap.Any("resp", resp),
			zap.Error(err),
			zap.Duration("cost", time.Since(start)),
		)
	}()

	resp, err = handler(ctx, req)
	return resp, err
}

// StreamServerRecoveryInterceptor recover
func StreamServerLogInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	var addr string
	if peer, ok := peer.FromContext(stream.Context()); ok {
		addr = peer.Addr.String()
	}
	var start = time.Now()
	defer func() {
		zlog.Info("grpc(*stream)",
			zap.String("service", path.Dir(info.FullMethod)[1:]),
			zap.String("method", path.Base(info.FullMethod)),
			zap.String("addr", addr),
			zap.Any("srv", srv),
			zap.Error(err),
			zap.Duration("cost", time.Since(start)),
		)
	}()

	return handler(srv, stream)
}
