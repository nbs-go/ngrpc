package ngrpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"strings"
	"time"
)

type RequestMetadata struct {
	RequestId string    `json:"requestId"`
	ClientIP  string    `json:"clientIP"`
	UserAgent string    `json:"userAgent"`
	StartedAt time.Time `json:"startedAt"`
}

func GetRequestMetadata(ctx context.Context) *RequestMetadata {
	val, ok := ctx.Value(ContextKeyRequestMetadata).(RequestMetadata)
	if !ok {
		return &RequestMetadata{}
	}
	return &val
}

func NewRequestMetadata(ctx context.Context, trustProxy string) RequestMetadata {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		// Ensure header value is initialized
		headers = make(metadata.MD)
	}

	// Get user-agent header captured from gateway
	userAgent, _ := GetHeaderValue(headers, HeaderGatewayUserAgent)

	return RequestMetadata{
		RequestId: getRequestId(headers),
		ClientIP:  getClientIP(headers, ctx, trustProxy),
		UserAgent: userAgent,
		StartedAt: time.Now(),
	}
}

func getClientIP(headers metadata.MD, ctx context.Context, trustProxy string) string {
	ip, ok := GetHeaderValue(headers, HeaderRealIP)
	if ok && ip != "" {
		return strings.TrimSpace(ip)
	}

	if !ok && trustProxy != "" {
		ip, _ = GetHeaderValue(headers, HeaderForwardedFor)
	}

	// IP is empty, then retrieve from peer
	p, _ := peer.FromContext(ctx)
	return p.Addr.String()
}

func getRequestId(headers metadata.MD) string {
	// Read request id from headers
	reqId, ok := GetHeaderValue(headers, HeaderRequestId)
	if ok && reqId != "" {
		return reqId
	}

	// Generate new request id with UUID
	u, err := uuid.NewUUID()
	if err != nil {
		panic(fmt.Errorf("ngrpc: unable to generate uuid for request id. Cause=%w", err))
	}
	return u.String()
}
