package middleware

import (
	"context"
	"path"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/zlog"
	jsoniter "github.com/json-iterator/go"
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
		zlog.Info("",
			zap.String("addr", addr),
			zap.Duration("cost", time.Since(start)),
			zap.String("service", path.Dir(info.FullMethod)[1:]),
			zap.String("method", path.Base(info.FullMethod)),
			zap.String("req", jsonFormat(req)),
			zap.String("resp", jsonFormat(resp)),
			zap.Error(err),
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
		zlog.Info("",
			zap.String("addr", addr),
			zap.Duration("cost", time.Since(start)),
			zap.String("service", path.Dir(info.FullMethod)[1:]),
			zap.String("method", path.Base(info.FullMethod)),
			zap.String("srv", jsonFormat(srv)),
			zap.Error(err),
		)
	}()

	return handler(srv, stream)
}

func jsonFormat(data interface{}) string {
	buf, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err == nil {
		return string(buf)
	}

	if a, ok := data.(StringAble); ok {
		return a.String()
	}

	return ""
}

// StringAble string
type StringAble interface {
	String() string
}
