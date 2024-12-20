package adapter

import (
	"bff/internal/management/domain"
	"context"
)

type ManagementService interface {
	RegistryApi(ctx context.Context, api domain.API) error
	GetServices(ctx context.Context) ([]string, error)
	GetEndpoints(ctx context.Context, serviceName string) (domain.API, error)
	RemoveService(ctx context.Context, serviceName string) error
	RemoveEndpoint(ctx context.Context, endpointName string) error
}
