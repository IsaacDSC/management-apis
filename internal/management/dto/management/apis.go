package management

import (
	"bff/internal/management/domain"
	"bff/internal/management/dto/postman"
)

type API struct {
	ServiceName string             `json:"service_name"`
	Endpoints   []postman.Endpoint `json:"endpoints"`
}

func ToDomain(api API) domain.API {
	endpoint := make([]domain.Endpoint, len(api.Endpoints))
	for i := range api.Endpoints {
		endpoint[i] = api.Endpoints[i].ToDomain()
	}

	return domain.API{
		ServiceName: api.ServiceName,
		Endpoints:   endpoint,
	}
}
