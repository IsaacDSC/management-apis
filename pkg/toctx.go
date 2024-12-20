package pkg

import (
	"context"
	"github.com/google/uuid"
)

func SetToCtx(ctx context.Context, key any, correlation uuid.UUID) context.Context {
	return context.WithValue(ctx, key, correlation)
}

func GetFromCtx(ctx context.Context, key any) uuid.UUID {
	if correlation, ok := ctx.Value(key).(uuid.UUID); ok {
		return correlation
	}
	return uuid.New()
}
