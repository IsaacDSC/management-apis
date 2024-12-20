package restapi

import (
	"bff/internal/infra/database/sqlc"
	"bff/internal/management/adapter"
	"context"
	"database/sql"
	"fmt"
	"regexp"
)

func GetRouters(ctx context.Context, db *sql.DB) (adapter.HttpAdapterHandler, error) {
	conn := sqlc.New(db)
	endpoints, err := conn.GetAllEndpoints(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting endpoints: %w", err)
	}

	routers := adapter.HttpAdapterHandler{}

	keyValues := map[string]string{}
	const prefix = "/api/v1/%s%s"
	re := regexp.MustCompile(`^/api/v1/[^/]+`)

	for _, endpoint := range endpoints {
		keyValues[endpoint.Name] = endpoint.Url
		result := re.ReplaceAllString(fmt.Sprintf(prefix, endpoint.Name, endpoint.Path), "/api/v1/{endpoint_name}")
		routers[fmt.Sprintf("%s %s", endpoint.Method, result)] = DefaultProxyHandler
	}

	if err := SaveCache(endpoints); err != nil {
		return nil, fmt.Errorf("error saving cache: %w", err)
	}

	return routers, nil
}
