package adapter

import (
	"bff/internal/management/domain"
	"context"
)

type ManagementRepository interface {
	Save(ctx context.Context, api domain.API) error
	GetEndpoints(ctx context.Context, serviceName string) (domain.API, error)
	GetServices(ctx context.Context) ([]string, error)
	RemoveService(ctx context.Context, serviceName string) error
	RemoveEndpoint(ctx context.Context, endpointName string) error
}
