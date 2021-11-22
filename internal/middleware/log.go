package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/zlog"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// UnaryServerLogInterceptor log 拦截
func UnaryServerLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	method := info.FullMethod[strings.LastIndex(info.FullMethod, "/")+1:]
	var addr string
	if peer, ok := peer.FromContext(ctx); ok {
		addr = peer.Addr.String()
	}
	reqBody := ""
	if str, ok := respToString(req); ok {
		reqBody = str
	}

	log := &logInfo{
		Method: method,
		Start:  time.Now().UnixNano(),
		IP:     addr,
		Req:    reqBody,
	}

	defer func() {
		cost := time.Since(time.Unix(0, log.Start))
		log.Cost = cost
		if str, ok := respToString(resp); ok {
			log.Resp = str
		}

		if e := recover(); e != nil {
			log.Code = -1
			log.ErrMsg = fmt.Sprintf("Panic: %v", e)
			err = fmt.Errorf("panic: %v", e)
		} else if err != nil {
			if s, ok := status.FromError(err); ok {
				log.Code = int64(s.Code())
				log.ErrMsg = s.Message()
			} else {
				log.Code = 0
				log.ErrMsg = err.Error()
			}
		}
		zlog.Info("", zap.String("grpc data", log.String()))
	}()

	resp, err = handler(ctx, req)
	return resp, err
}

func respToString(resp interface{}) (string, bool) {
	data, err := json.Marshal(resp)
	if err == nil {
		return string(data), true
	}

	if a, ok := resp.(StringAble); ok {
		return a.String(), true
	}

	return "", false
}

type logInfo struct {
	Method string
	Start  int64
	Cost   time.Duration
	IP     string
	Req    string
	Resp   string
	Code   int64
	ErrMsg string
}

func (l *logInfo) String() string {
	buf, err := json.Marshal(l)
	if err != nil {
		return `{"errMsg": "Missing log"}`
	}
	return string(buf)
}

// StringAble string
type StringAble interface {
	String() string
}
