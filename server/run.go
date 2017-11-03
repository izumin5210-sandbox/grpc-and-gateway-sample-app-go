package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/creasty/apperrors"
	"github.com/soheilhy/cmux"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/system"
)

// Run initializes and start application server
func Run(c context.Context, appC *system.AppContext) error {
	lis, err := net.Listen("tcp", appC.Host)
	if err != nil {
		return apperrors.WithMessage(err, fmt.Sprintf("failed to listen %s", appC.Host))
	}
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := cmux.New(lis)

	{
		// Starts gRPC server
		s := grpcServer(c, appC)
		go s.Serve(mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc")))
	}

	{
		// Starts grpc-gateway server
		s, err := gatewayServer(c, appC)
		if err != nil {
			return apperrors.WithMessage(err, "failed to start grpc-gateway server")
		}
		go http.Serve(mux.Match(cmux.HTTP2(), cmux.HTTP1()), s)
	}

	return mux.Serve()
}
