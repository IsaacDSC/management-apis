package containers

import (
	"bff/internal/management/adapter"
	"bff/internal/management/infra/gateway"
)

type GatewaysContainer struct {
	Proxy adapter.RequestApiAdapter
}

func NewGatewaysContainer() *GatewaysContainer {
	return &GatewaysContainer{
		Proxy: gateway.NewRequestApi(),
	}
}
