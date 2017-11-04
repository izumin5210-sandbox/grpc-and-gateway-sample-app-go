package interceptor

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/creasty/apperrors"
	"github.com/getsentry/raven-go"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/izumin5210/grpc-errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

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

var jsonPbMarshaller = &jsonpb.Marshaler{}

var notWrappedErrorHandler = grpcerrors.WithNotWrappedErrorHandler(func(_ context.Context, err error) error {
	return system.ErrorUnknown.WithReport(err)
})

var statusCodeMapper = grpcerrors.WithStatusCodeMapper(func(code int) codes.Code {
	grpcCode, ok := grpcCodeBySystemCode[system.ErrorCode(code)]
	if !ok {
		return codes.Unknown
	}
	return grpcCode
})

var unaryServerErrorHandlers = []grpcerrors.UnaryServerErrorHandler{
	notWrappedErrorHandler,
	grpcerrors.WithUnaryServerReportableErrorHandler(func(c context.Context, req interface{}, info *grpc.UnaryServerInfo, err *apperrors.Error) error {
		st := &raven.Stacktrace{}
		for _, t := range err.StackTrace {
			f := &raven.StacktraceFrame{
				Filename: t.File,
				Function: t.Func,
				Lineno:   int(t.Line),
			}
			// Sentry wants the frames with the oldest first, so reverse them
			st.Frames = append([]*raven.StacktraceFrame{f}, st.Frames...)
		}
		pckt := raven.NewPacket(err.Error(), st)
		if p, ok := req.(proto.Message); ok {
			jbuf := &bytes.Buffer{}
			if err := jsonPbMarshaller.Marshal(jbuf, p); err == nil {
				var data map[string]interface{}
				if err := json.Unmarshal(jbuf.Bytes(), &data); err == nil {
					pckt.Extra["body"] = data
				}
			}
		}
		if md, ok := metadata.FromIncomingContext(c); ok {
			pckt.Extra["metadata"] = md
		}
		pckt.Extra["context_tags"] = grpc_ctxtags.Extract(c)

		raven.Capture(pckt, map[string]string{
			"method": info.FullMethod,
		})
		return err
	}),
	statusCodeMapper,
}

var streamServerErrorHandlers = []grpcerrors.StreamServerErrorHandler{
	notWrappedErrorHandler,
	grpcerrors.WithReportableErrorHandler(func(_ context.Context, err *apperrors.Error) error {
		// TODO: wanna report errors with request contexts...
		raven.CaptureError(err, map[string]string{})
		return err
	}),
	statusCodeMapper,
}

func errorsStreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpcerrors.StreamServerInterceptor(streamServerErrorHandlers...)
}

func errorsUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpcerrors.UnaryServerInterceptor(unaryServerErrorHandlers...)
}
