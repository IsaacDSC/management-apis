package pkg

import (
	"context"
	"github.com/google/uuid"
)

type correlationStruct struct{}

var correlationIDKey = correlationStruct{}

func SetCorrelationID(ctx context.Context, correlation uuid.UUID) context.Context {
	return SetToCtx(ctx, correlationIDKey, correlation)
}

func GetCorrelationID(ctx context.Context) uuid.UUID {
	return GetFromCtx(ctx, correlationIDKey)
}
