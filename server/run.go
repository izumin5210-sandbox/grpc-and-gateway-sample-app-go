package server

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run initializes and start application server
func Run(c context.Context) error {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	reflection.Register(s)
	return s.Serve(lis)
}
