package server

import (
	"context"
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
)

// Run initializes and start application server
func Run(c context.Context) error {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := cmux.New(lis)

	{
		// Starts gRPC server
		s := grpcServer(c)
		go s.Serve(mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc")))
	}

	{
		// Starts grpc-gateway server
		s, err := gatewayServer(c)
		if err != nil {
			return err
		}
		go http.Serve(mux.Match(cmux.HTTP2(), cmux.HTTP1()), s)
	}

	return mux.Serve()
}
