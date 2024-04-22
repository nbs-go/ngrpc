package ngrpc

import (
	"errors"
	"fmt"
	"github.com/nbs-go/errx"
	"google.golang.org/grpc/codes"
)

func NewErrorDetails(err error, withSource bool) (codes.Code, *ErrorDetails) {
	var hErr *errx.Error
	ok := errors.As(err, &hErr)
	if !ok {
		// If assert type fail, create wrap error to an internal error
		hErr = errx.InternalError().Wrap(err)
	}

	// Get grpc status
	errMeta := hErr.Metadata()
	grpcStatus, ok := errMeta[KeyGrpcStatus].(codes.Code)
	if !ok {
		grpcStatus = codes.Internal
	}

	// Compose details
	details := &ErrorDetails{
		Code:    hErr.Code(),
		Message: hErr.Message(),
	}

	if !withSource {
		return grpcStatus, details
	}

	// Compose source message
	sourceMsg := ""
	if sourceErr := errors.Unwrap(hErr); sourceErr != nil {
		sourceMsg = sourceErr.Error()
	} else {
		sourceMsg = hErr.Message()
	}

	// Compose metadata
	rawMetadata, _ := errMeta[KeyErrorMetadata].(map[string]interface{})
	var metadata map[string]string
	if rawMetadata != nil {
		metadata = make(map[string]string)
		for k, v := range rawMetadata {
			metadata[k] = fmt.Sprintf("%v", v)
		}
	}

	// Compose source error
	details.Source = &ErrorDetails_Source{
		Message:  sourceMsg,
		Traces:   hErr.Traces(),
		Metadata: metadata,
	}

	return grpcStatus, details
}
