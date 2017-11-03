package server

import (
	"context"
	"net/http"

	"github.com/creasty/apperrors"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/api"
	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/server/interceptor"
	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/system"
)

func gatewayServer(c context.Context, appC *system.AppContext) (http.Handler, error) {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		interceptor.WithStreamClientInterceptor(appC),
		interceptor.WithUnaryClientInterceptor(appC),
	}

	var err error

	err = api.RegisterProfileServiceHandlerFromEndpoint(c, mux, appC.Host, opts)
	if err != nil {
		return nil, apperrors.WithMessage(err, "failed to register ProfileServiceServer handler")
	}

	return mux, nil
}
