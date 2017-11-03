package interceptor

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/system"
)

// WithStreamServerInterceptor returns composed StreamServerInterceptor
func WithStreamServerInterceptor(c *system.AppContext) grpc.ServerOption {
	interceptors := []grpc.StreamServerInterceptor{}
	interceptors = append(interceptors,
		errorsStreamServerInterceptor(),
		recoveryStreamServerInterceptor(),
		ctxtagsStreamServerInterceptor(),
		zapStreamServerInterceptor(c),
	)

	return grpc_middleware.WithStreamServerChain(interceptors...)
}

// WithUnaryServerInterceptor returns composed UnaryServerInterceptor
func WithUnaryServerInterceptor(c *system.AppContext) grpc.ServerOption {
	interceptors := []grpc.UnaryServerInterceptor{}
	interceptors = append(interceptors,
		errorsUnaryServerInterceptor(),
		recoveryUnaryServerInterceptor(),
		ctxtagsUnaryServerInterceptor(),
		zapUnaryServerInterceptor(c),
	)

	return grpc_middleware.WithUnaryServerChain(interceptors...)
}

// WithStreamClientInterceptor returns composed StreamClientInterceptor
func WithStreamClientInterceptor(c *system.AppContext) grpc.DialOption {
	interceptors := []grpc.StreamClientInterceptor{}
	interceptors = append(interceptors,
		zapStreamClientInterceptor(c),
	)
	return grpc.WithStreamInterceptor(
		grpc_middleware.ChainStreamClient(interceptors...),
	)
}

// WithUnaryClientInterceptor returns composed UnaryClientInterceptor
func WithUnaryClientInterceptor(c *system.AppContext) grpc.DialOption {
	interceptors := []grpc.UnaryClientInterceptor{}
	interceptors = append(interceptors,
		zapUnaryClientInterceptor(c),
	)
	return grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(interceptors...),
	)
}
