package postman

import (
	"bff/internal/management/domain"
	"github.com/google/uuid"
)

type Endpoint struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Method      string              `json:"method"`
	URL         string              `json:"url"`
	Path        string              `json:"path"`
	Headers     map[string][]string `json:"headers"`
	Body        map[string]any      `json:"body"`
}

func (e Endpoint) ToDomain() domain.Endpoint {
	return domain.Endpoint{
		ID:          uuid.New(),
		Name:        e.Name,
		Description: e.Description,
		Method:      e.Method,
		URL:         e.URL,
		Path:        e.Path,
		Headers:     e.Headers,
		Body:        e.Body,
	}
}
