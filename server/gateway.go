package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/api"
)

func gatewayServer(c context.Context) (http.Handler, error) {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	var err error

	err = api.RegisterProfileServiceHandlerFromEndpoint(c, mux, ":3000", opts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
