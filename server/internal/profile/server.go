package profile

import (
	"google.golang.org/grpc"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/api"
)

type server struct {
}

// Register sets ProfileService server implementation to given grpc.Server
func Register(s *grpc.Server) {
	srv := &server{}
	api.RegisterProfileServiceServer(s, srv)
}
