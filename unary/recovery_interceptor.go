package unary

import (
	"context"
	"github.com/nbs-go/ngrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
)

type RecoveryInterceptor struct {
	handler  ngrpc.RequestHandler
	debugErr bool
}

func NewRecoveryInterceptor(handler ngrpc.RequestHandler, debugErr bool) *RecoveryInterceptor {
	if handler == nil {
		handler = &noHandler{}
	}

	return &RecoveryInterceptor{
		handler:  handler,
		debugErr: debugErr,
	}
}

func (i *RecoveryInterceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
	// Get method name
	panicked := true

	// Recover from panic
	defer func() {
		if r := recover(); r != nil || panicked {
			i.handler.OnPanic(ctx, info.FullMethod, r, debug.Stack(), req)
			err = status.Errorf(codes.Internal, "%v", r)
		}
	}()

	// Handle request
	resp, err := handler(ctx, req)
	if err != nil {
		err = i.wrapError(ctx, err)
	}

	// Get metadata
	reqMeta := ngrpc.GetRequestMetadata(ctx)
	i.handler.OnRequestHandled(ctx, info.FullMethod, getGrpcStatus(err), reqMeta)

	panicked = false
	return resp, err
}

func getGrpcStatus(err error) codes.Code {
	if err == nil {
		return codes.OK
	}

	sErr, ok := status.FromError(err)
	if !ok {
		return codes.Internal
	}

	return sErr.Code()
}

func (i *RecoveryInterceptor) wrapError(ctx context.Context, err error) error {
	if i.handler.IsCanceled(err) {
		err = ngrpc.CancelError.Wrap(err)
	}

	// Get error details for grpc
	grpcStatus, errDetails := ngrpc.NewErrorDetails(err, i.debugErr)

	// Wrap to grpc error
	st := status.New(grpcStatus, errDetails.Message)
	result, gErr := st.WithDetails(errDetails)
	if gErr != nil {
		panic(gErr)
	}

	i.handler.OnRequestError(ctx, err)

	return result.Err()
}

// noHandler implements ngrpc.RequestHandler to do nothing. This handler will be used as fallback handler
type noHandler struct{}

func (n *noHandler) IsCanceled(_ error) bool {
	return false
}

func (n *noHandler) OnRequestError(_ context.Context, _ error) {
	// Do nothing
}

func (n *noHandler) OnRequestHandled(_ context.Context, _ string, _ codes.Code, _ *ngrpc.RequestMetadata) {
	// Do nothing
}

func (n *noHandler) OnPanic(_ context.Context, _ string, _ any, _ []byte, _ any) {
	// Do nothing
}
