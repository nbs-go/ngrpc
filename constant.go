package ngrpc

const (
	pkgNamespace = "ngrpc"
)

const (
	HeaderRealIP           = "x-real-ip"
	HeaderForwardedFor     = "x-forwarded-for"
	HeaderSubjectId        = "x-subject-id"
	HeaderSubjectFullName  = "x-subject-name"
	HeaderSubjectRole      = "x-subject-role"
	HeaderGatewayUserAgent = "x-gateway-user-agent"
	HeaderRequestId        = "x-request-id"
	HeaderHttpStatus       = "x-http-status"
	HeaderUserAgent        = "user-agent"
)

type ContextKey string

const (
	ContextKeyBasicAuth       ContextKey = "basic-auth"
	ContextKeyBearerToken     ContextKey = "bearer-token"
	ContextKeyRequestMetadata ContextKey = "request-metadata"
	ContextKeySubject         ContextKey = "subject"
)

const (
	KeyGrpcStatus      = "grpcStatus"
	KeyOverrideMessage = "message"
	KeyErrorMetadata   = "errorMetadata"
)

const (
	ResponseCodeSuccess    = "OK"
	ResponseMessageSuccess = "Success"
)

const (
	GrpcMetadataHttpStatus  = "Grpc-Metadata-X-Http-Status"
	GrpcMetadataContentType = "Grpc-Metadata-Content-Type"
	GrpcMetadataRequestId   = "Grpc-Metadata-X-Request-Id"
)
