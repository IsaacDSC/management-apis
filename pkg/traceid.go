package pkg

import (
	"context"
	"github.com/google/uuid"
)

type traceIDStruct struct{}

var traceIDKey = traceIDStruct{}

func SetTraceID(ctx context.Context, correlation uuid.UUID) context.Context {
	return SetToCtx(ctx, traceIDKey, correlation)
}

func GetTraceID(ctx context.Context) uuid.UUID {
	return GetFromCtx(ctx, traceIDKey)
}
