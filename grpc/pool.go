package dev_tempalte_grpc

import (
	"github.com/farwydi/go-dev-template"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/processout/grpc-go-pool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func MakePool(addr string) (pool *grpcpool.Pool) {
	var err error
	factory := func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(addr, []grpc.DialOption{
			grpc.WithUnaryInterceptor(grpc_zap.UnaryClientInterceptor(dev_tempalte.ZapLogger)),
			grpc.WithInsecure(),
		}...)
		if err != nil {
			dev_tempalte.ZapLogger.Error("Failed to start gRPC connection",
				zap.Error(err),
				zap.String("addr", addr),
			)
		}
		return conn, err
	}
	pool, err = grpcpool.New(factory, 4, 8, time.Minute)
	if err != nil {
		dev_tempalte.ZapLogger.Panic("Failed to create gRPC pool",
			zap.Error(err),
			zap.String("addr", addr),
		)
	}

	return
}
