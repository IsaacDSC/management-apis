package containers

import (
	"bff/internal/management/adapter"
	"bff/internal/management/service"
)

type ServicesContainer struct {
	PostmanCollection adapter.PostmanCollectionService
	Management        adapter.ManagementService
}

func NewServicesContainer(gc *GatewaysContainer, rc *RepositoriesContainer) ServicesContainer {
	return ServicesContainer{
		PostmanCollection: service.NewPostmanCollection(rc.PostmanCollection),
		Management:        service.NewService(gc.Proxy, rc.Management),
	}
}
