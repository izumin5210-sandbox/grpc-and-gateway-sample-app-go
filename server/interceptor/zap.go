package interceptor

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/system"
)

func zapStreamServerInterceptor(ctx *system.AppContext) grpc.StreamServerInterceptor {
	return grpc_zap.PayloadStreamServerInterceptor(
		ctx.Logger,
		func(context.Context, string, interface{}) bool { return true },
	)
}

func zapUnaryServerInterceptor(ctx *system.AppContext) grpc.UnaryServerInterceptor {
	return grpc_zap.PayloadUnaryServerInterceptor(
		ctx.Logger,
		func(context.Context, string, interface{}) bool { return true },
	)
}

func zapStreamClientInterceptor(ctx *system.AppContext) grpc.StreamClientInterceptor {
	return grpc_zap.PayloadStreamClientInterceptor(
		ctx.Logger,
		func(context.Context, string) bool { return true },
	)
}

func zapUnaryClientInterceptor(ctx *system.AppContext) grpc.UnaryClientInterceptor {
	return grpc_zap.PayloadUnaryClientInterceptor(
		ctx.Logger,
		func(context.Context, string) bool { return true },
	)
}
