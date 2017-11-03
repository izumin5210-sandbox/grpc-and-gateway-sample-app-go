package interceptor

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

var ctxtagsOptions = []grpc_ctxtags.Option{
	grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
}

func ctxtagsStreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc_ctxtags.StreamServerInterceptor(ctxtagsOptions...)
}

func ctxtagsUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_ctxtags.UnaryServerInterceptor(ctxtagsOptions...)
}
