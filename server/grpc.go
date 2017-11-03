package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/server/internal/profile"
	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/system"
)

func grpcServer(c context.Context, appC *system.AppContext) *grpc.Server {
	s := grpc.NewServer()
	profile.Register(s)
	reflection.Register(s)
	return s
}
