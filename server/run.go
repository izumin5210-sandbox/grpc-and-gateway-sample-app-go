package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/creasty/apperrors"
	"github.com/soheilhy/cmux"
)

// Run initializes and start application server
func Run(c context.Context) error {
	host := ":3000"
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return apperrors.WithMessage(err, fmt.Sprintf("failed to listen %s", host))
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
			return apperrors.WithMessage(err, "failed to start grpc-gateway server")
		}
		go http.Serve(mux.Match(cmux.HTTP2(), cmux.HTTP1()), s)
	}

	return mux.Serve()
}
