package gateway

import (
	"context"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nbs-go/ngrpc"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
)

func HandleCaptureMetadata(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set request id
		reqId, err := uuid.NewUUID()
		if err != nil {
			panic(err)
		}

		// Forward request id to grpc
		r.Header.Set(ngrpc.HeaderRequestId, reqId.String())

		// Set request id to writer as fallback request id
		w.Header().Set(ngrpc.HeaderRequestId, "gw-"+reqId.String())

		// Capture User-Agent value and set to Gateway User-Agent
		ua := r.Header.Get(ngrpc.HeaderUserAgent)
		r.Header.Set(ngrpc.HeaderGatewayUserAgent, ua)

		// Capture X-Real-IP Value
		clientIp := r.Header.Get(ngrpc.HeaderRealIP)
		r.Header.Set(ngrpc.HeaderRealIP, clientIp)

		// Continue request to grpc
		h.ServeHTTP(w, r)
	}
}

func HandleForwardMetadataResponse(ctx context.Context, w http.ResponseWriter, _ proto.Message) error {
	// Get all grpc metadata / response headers
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// Set request id to response
	vals := md.HeaderMD.Get(ngrpc.HeaderRequestId)
	if len(vals) > 0 {
		w.Header().Set(ngrpc.HeaderRequestId, vals[0])
		// clean up headers from response
		delete(w.Header(), ngrpc.GrpcMetadataRequestId)
	}

	// Override http status code
	vals = md.HeaderMD.Get(ngrpc.HeaderHttpStatus)
	if len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, ngrpc.HeaderHttpStatus)
		delete(w.Header(), ngrpc.GrpcMetadataHttpStatus)
		w.WriteHeader(code)
	}

	// Clean up unnecessary header
	delete(w.Header(), ngrpc.GrpcMetadataContentType)

	return nil
}
