package unary

import (
	"context"
	"encoding/base64"
	"github.com/nbs-go/ngrpc"
	logContext "github.com/nbs-go/nlogger/v2/context"
	logOption "github.com/nbs-go/nlogger/v2/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

type RequestInterceptor struct {
	trustProxy string
}

func NewRequestInterceptor(trustProxy string) *RequestInterceptor {
	return &RequestInterceptor{trustProxy: trustProxy}
}

func (i *RequestInterceptor) InterceptAuthorization(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Get authorization header value
	authType, authValue := getAuthorizationHeaderValue(ctx, "authorization")
	if authType == "" || authValue == "" {
		return handler(ctx, req)
	}

	switch authType {
	case "Basic":
		ctx = setBasicAuth(ctx, authValue)
	case "Bearer":
		ctx = context.WithValue(ctx, ngrpc.ContextKeyBearerToken, authValue)
	}

	return handler(ctx, req)
}

func (i *RequestInterceptor) InterceptMetadata(ctx context.Context, req interface{}, srv *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Parse request metadata from context header
	reqMeta := ngrpc.NewRequestMetadata(ctx, i.trustProxy)

	// Set request metadata
	ctx = context.WithValue(ctx, ngrpc.ContextKeyRequestMetadata, reqMeta)

	// Inject request id to context for logging
	reqId := reqMeta.RequestId
	ctx = context.WithValue(ctx, logContext.RequestIdKey, reqId)

	// Set request id to response headers
	err := grpc.SetHeader(ctx, metadata.Pairs(ngrpc.HeaderRequestId, reqId))
	if err != nil {
		log.Warnf("Unable to set request id to context. RequestId=%s Error=\"%s\"", reqId, err)
	}

	log.Trace("Start Handling request. FullMethod=%s ClientIP=%s UserAgent=%s Payload=%+v", logOption.Format(srv.FullMethod, reqMeta.ClientIP, reqMeta.UserAgent, req), logOption.Context(ctx))
	return handler(ctx, req)
}

func (i *RequestInterceptor) InterceptSubject(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	subject := getSubjectFromHeader(ctx)
	ctx = context.WithValue(ctx, ngrpc.ContextKeySubject, subject)
	return handler(ctx, req)
}

func getSubjectFromHeader(ctx context.Context) ngrpc.Subject {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ngrpc.AnonymousUser
	}

	// Get values
	id, _ := ngrpc.GetHeaderValue(headers, ngrpc.HeaderSubjectId)
	fullName, _ := ngrpc.GetHeaderValue(headers, ngrpc.HeaderSubjectFullName)
	role, _ := ngrpc.GetHeaderValue(headers, ngrpc.HeaderSubjectRole)

	if id == "" && fullName == "" && role == "" {
		return ngrpc.AnonymousUser
	}

	return ngrpc.Subject{
		Id:       id,
		FullName: fullName,
		Role:     role,
	}
}

func getAuthorizationHeaderValue(ctx context.Context, header string) (authType string, value string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ""
	}

	authHeaders, ok := md[header]
	if !ok {
		return "", ""
	}

	if len(authHeaders) != 1 {
		return "", ""
	}

	// Parse
	tmp := strings.SplitN(authHeaders[0], " ", 2)
	if len(tmp) != 2 {
		return "", ""
	}

	return tmp[0], tmp[1]
}

func setBasicAuth(ctx context.Context, authValue string) context.Context {
	c, err := base64.StdEncoding.DecodeString(authValue)
	if err != nil {
		return ctx
	}

	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return ctx
	}

	// Get username and password
	username, password := cs[:s], cs[s+1:]

	// Set to context
	return context.WithValue(ctx, ngrpc.ContextKeyBasicAuth, ngrpc.BasicAuth{username, password})
}
