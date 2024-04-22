package ngrpc

import (
	"context"
)

func GetBearerToken(ctx context.Context) string {
	val, ok := ctx.Value(ContextKeyBearerToken).(string)
	if !ok {
		return ""
	}
	return val
}
