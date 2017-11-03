package interceptor

import (
	"fmt"

	"github.com/creasty/apperrors"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

var recoveryHandler = grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
	if err, ok := p.(error); ok {
		return apperrors.Wrap(err)
	}
	return apperrors.New(fmt.Sprint(p))
})

func recoveryStreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc_recovery.StreamServerInterceptor(recoveryHandler)
}

func recoveryUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor(recoveryHandler)
}
