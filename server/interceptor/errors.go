package interceptor

import (
	"context"

	"github.com/creasty/apperrors"
	"github.com/izumin5210/grpc-errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/system"
)

var grpcCodeBySystemCode = map[system.ErrorCode]codes.Code{
	system.ErrorUnknown:         codes.Unknown,
	system.ErrorInvalidArgument: codes.InvalidArgument,
	system.ErrorUnauthorized:    codes.Unauthenticated,
	system.ErrorForbbiden:       codes.PermissionDenied,
	system.ErrorNotFound:        codes.NotFound,
	system.ErrorFailedToReadDB:  codes.Internal,
	system.ErrorFailedToWriteDB: codes.Internal,
}

var errorHandleFuncs = []grpcerrors.ErrorHandlerFunc{
	grpcerrors.WithNotWrappedErrorHandler(func(_ context.Context, err error) error {
		return system.ErrorUnknown.WithReport(err)
	}),
	grpcerrors.WithReportableErrorHandler(func(_ context.Context, err *apperrors.Error) error {
		// TODO: Should send error reports
		return err
	}),
	grpcerrors.WithStatusCodeMapper(func(code int) codes.Code {
		grpcCode, ok := grpcCodeBySystemCode[system.ErrorCode(code)]
		if !ok {
			return codes.Unknown
		}
		return grpcCode
	}),
}

func errorsStreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpcerrors.StreamServerInterceptor(errorHandleFuncs...)
}

func errorsUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpcerrors.UnaryServerInterceptor(errorHandleFuncs...)
}
