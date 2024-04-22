package gateway

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nbs-go/ngrpc"
	logOption "github.com/nbs-go/nlogger/v2/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func NewGrpcErrorHandler(debugErr bool) runtime.ErrorHandlerFunc {
	return func(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
		reqId := getRequestId(ctx, w)
		httpStatus, errBody := composeErrorResponse(err, debugErr, reqId)

		// Set headers
		w.Header().Set(ngrpc.HeaderContentType, marshaler.ContentType(&errBody))
		w.Header().Set(ngrpc.HeaderRequestId, reqId)
		w.WriteHeader(httpStatus)

		// Write body to json
		jErr := json.NewEncoder(w).Encode(errBody)

		if jErr != nil {
			log.Warnf("Unable to write json error to response writer. Error=%s", jErr)

			// Fallback to string error
			_, wErr := w.Write([]byte(internalErrorJson))
			if wErr != nil {
				log.Warnf("Unable to write string error to response writer. Error=%s", wErr)
			}
		}
	}
}

func composeErrorResponse(err error, debug bool, reqId string) (httpStatus int, body Response) {
	// Get error details
	gErr := status.Convert(err)
	errCode := gErr.Code()
	log.Error("GrpcStatus=%d GrpcMessage=\"%s\"", logOption.Format(gErr.Code(), gErr.Message()),
		logOption.Error(err), logOption.AddMetadata("requestId", reqId))

	// If no error details, then return internal error
	errDetails := gErr.Details()
	if len(errDetails) == 0 {
		switch errCode {
		case codes.InvalidArgument:
			return http.StatusBadRequest, BadRequestError
		case codes.NotFound, codes.Unimplemented:
			return http.StatusNotFound, NotFoundError
		}
		return http.StatusInternalServerError, InternalError
	}

	// Cast error to ErrorResult
	aErr, ok := errDetails[0].(*ngrpc.ErrorDetails)
	if !ok {
		return http.StatusInternalServerError, InternalError
	}

	errBody := Response{
		Success: false,
		Code:    aErr.Code,
		Message: aErr.Message,
		Data:    nil,
	}

	// Write error cause on debug mode
	if sErr := aErr.Source; debug && sErr != nil {
		errBody.Data = &ErrorData{
			Debug: &ErrorDebugData{
				Message:  sErr.Message,
				Traces:   sErr.Traces,
				Metadata: sErr.Metadata,
			},
		}
	}

	return runtime.HTTPStatusFromCode(gErr.Code()), errBody
}

// getRequestId retrieve request id from grpc and fallback to gateway request id if not found
func getRequestId(ctx context.Context, w http.ResponseWriter) string {
	// Get request id from
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if ok {
		// Set request id to response
		vals := md.HeaderMD.Get(ngrpc.HeaderRequestId)
		if len(vals) > 0 && vals[0] != "" {
			return vals[0]
		}
	}
	return w.Header().Get(ngrpc.HeaderRequestId)
}
