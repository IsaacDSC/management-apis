package repository

import (
	"bff/internal/infra/database/sqlc"
	"bff/internal/management/domain"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

type Management struct {
	db *sql.DB
}

func NewManagement(db *sql.DB) *Management {
	return &Management{db: db}
}

func (m Management) Save(ctx context.Context, api domain.API) error {
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("error on begin transaction: %w", err)
	}

	orm := sqlc.New(tx)
	defer tx.Rollback()

	for i := range api.Endpoints {
		headers, err := json.Marshal(api.Endpoints[i].Headers)
		if err != nil {
			return fmt.Errorf("error on marshal headers: %w", err)
		}

		body, err := json.Marshal(api.Endpoints[i].Body)
		if err != nil {
			return fmt.Errorf("error on marshal body: %w", err)
		}

		if err := orm.CreateOrUpdate(ctx, sqlc.CreateOrUpdateParams{
			ServiceName:  api.ServiceName,
			Name:         fmt.Sprintf("%s_%s", api.ServiceName, api.Endpoints[i].Name),
			Description:  api.Endpoints[i].Description,
			Method:       api.Endpoints[i].Method,
			Url:          api.Endpoints[i].URL,
			Path:         api.Endpoints[i].Path,
			Headers:      string(headers),
			Body:         string(body),
			SensitiveApi: false,
			Active:       true,
		}); err != nil {
			return fmt.Errorf("error on create or update endpoint: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error on commit transaction: %w", err)
	}

	return nil
}

func (m Management) GetEndpoints(ctx context.Context, serviceName string) (domain.API, error) {
	orm := sqlc.New(m.db)
	endpoints, err := orm.GetEndpoints(ctx, serviceName)
	if err != nil {
		return domain.API{}, fmt.Errorf("error on get endpoints: %w", err)
	}

	api := domain.API{
		ServiceName: serviceName,
		Endpoints:   make([]domain.Endpoint, len(endpoints)),
	}

	for i, endpoint := range endpoints {
		var body map[string]any
		if endpoint.Body == "" {
			if err := json.Unmarshal([]byte(endpoint.Body), &body); err != nil {
				return domain.API{}, fmt.Errorf("error on unmarshal body: %w", err)
			}
		}

		api.Endpoints[i] = domain.Endpoint{
			ID:          endpoint.ID,
			Name:        endpoint.Name,
			Description: endpoint.Description,
			Method:      endpoint.Method,
			URL:         endpoint.Url,
			Path:        endpoint.Path,
			Body:        body,
		}

	}

	return api, nil
}

func (m Management) GetServices(ctx context.Context) ([]string, error) {
	orm := sqlc.New(m.db)
	services, err := orm.GetServices(ctx)
	if err != nil {
		return nil, fmt.Errorf("error on get services: %w", err)
	}

	return services, nil
}

func (m Management) RemoveService(ctx context.Context, serviceName string) error {
	orm := sqlc.New(m.db)
	return orm.RemoveService(ctx, serviceName)
}

func (m Management) RemoveEndpoint(ctx context.Context, endpointName string) error {
	orm := sqlc.New(m.db)
	return orm.RemoveEndpoint(ctx, endpointName)
}
