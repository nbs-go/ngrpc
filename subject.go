package ngrpc

import (
	"context"
)

type Subject struct {
	Id       string `json:"id"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
}

func GetSubject(ctx context.Context) Subject {
	val, ok := ctx.Value(ContextKeySubject).(Subject)
	if !ok {
		return AnonymousUser
	}
	return val
}

func SetSubject(ctx context.Context, id, fullName, role string) context.Context {
	return context.WithValue(ctx, ContextKeySubject, Subject{
		Id:       id,
		FullName: fullName,
		Role:     role,
	})
}

var AnonymousUser = Subject{
	Id:       "ANON",
	FullName: "Anonymous User",
	Role:     "USER",
}
