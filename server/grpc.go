package server

import (
	"context"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/server/internal/profile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func grpcServer(c context.Context) *grpc.Server {
	s := grpc.NewServer()
	profile.Register(s)
	reflection.Register(s)
	return s
}
