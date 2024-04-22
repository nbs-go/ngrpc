package ngrpc

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GetHeaderValue(headers metadata.MD, key string) (string, bool) {
	values := headers.Get(key)

	if len(values) == 0 {
		return "", false
	}

	if values[0] == "" {
		return "", false
	}

	return values[0], true
}

func GetHeaderFromContext(ctx context.Context, key string) (string, bool) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	return GetHeaderValue(headers, key)
}
