package ngrpc

import (
	"context"
)

type BasicAuth struct {
	Username string
	Password string
}

func GetBasicAuth(ctx context.Context) *BasicAuth {
	val, ok := ctx.Value(ContextKeyBasicAuth).(BasicAuth)
	if !ok {
		return nil
	}
	return &val
}
