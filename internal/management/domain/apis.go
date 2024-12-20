package domain

type API struct {
	ServiceName string
	Endpoints   []Endpoint
}

func NewAPI(serviceName string, endpoints []Endpoint) API {
	return API{ServiceName: serviceName, Endpoints: endpoints}
}
