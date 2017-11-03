package server

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/server/interceptor"
	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/server/internal/profile"
	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/system"
)

func grpcServer(c context.Context, appC *system.AppContext) *grpc.Server {
	grpc_zap.ReplaceGrpcLogger(appC.Logger)
	s := grpc.NewServer(
		interceptor.WithStreamServerInterceptor(appC),
		interceptor.WithUnaryServerInterceptor(appC),
	)
	profile.Register(s)
	reflection.Register(s)
	return s
}
