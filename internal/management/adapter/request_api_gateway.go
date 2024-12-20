package adapter

import (
	"bff/internal/management/domain"
	"context"
)

type RequestApiAdapter interface {
	Request(ctx context.Context, item domain.Endpoint) (map[string]any, error)
}
