package ngrpc

import (
	"context"
	"google.golang.org/grpc/codes"
)

type RequestHandler interface {
	IsCanceled(err error) bool
	OnPanic(ctx context.Context, fullMethod string, errValue any, stack []byte, req any)
	OnRequestError(ctx context.Context, err error)
	OnRequestHandled(ctx context.Context, fullMethod string, status codes.Code, reqMeta *RequestMetadata)
}
