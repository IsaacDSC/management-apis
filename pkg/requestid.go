package pkg

import (
	"context"
	"github.com/google/uuid"
)

type requestIDStruct struct{}

var requestIDKey = requestIDStruct{}

func SetRequestID(ctx context.Context, correlation uuid.UUID) context.Context {
	return SetToCtx(ctx, requestIDKey, correlation)
}

func GetRequestID(ctx context.Context) uuid.UUID {
	return GetFromCtx(ctx, requestIDKey)
}
