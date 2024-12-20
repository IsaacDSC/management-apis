package postman

import (
	"bff/internal/management/domain"
	"bff/internal/management/dto/management"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostmanCollection(t *testing.T) {
	ID := uuid.New()
	mockHost := "http://example.com"
	envMapper := map[string]string{"host": mockHost}
	mockPath := "/contenthub/sdk/teste1"

	tests := []struct {
		name           string
		collection     management.CollectionDto
		expectedDomain []domain.Endpoint
		expectedError  bool
	}{
		{
			name: "Valid collection file",
			expectedDomain: []domain.Endpoint{
				{
					ID:          ID,
					Name:        "Test Endpoint",
					Description: "",
					Method:      "GET",
					URL:         fmt.Sprintf("%s/contenthub/sdk/teste1", mockHost),
					Headers:     nil,
					Body:        nil,
				},
			},
			collection: management.CollectionDto{
				Info: management.Info{
					PostmanID:  ID.String(),
					Name:       "Test CollectionDto",
					Schema:     "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
					ExporterID: "67890",
				},
				Item: []management.PostmanItem{
					{
						Name: "Test Endpoint",
						Request: management.PostmanRequest{
							Method: "GET",
							URL: management.URL{
								Raw: management.KeyValue(fmt.Sprintf("{{host}}%s", mockPath)),
							},
						},
					},
				},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := tt.collection.ToDomain(envMapper)
			assert.NoError(t, err)
			ignoreValidationID(d, ID)

			assert.Equal(t, tt.expectedDomain, d)
		})
	}
}

func ignoreValidationID(d []domain.Endpoint, ID uuid.UUID) {
	for i := range d {
		d[i].ID = ID
	}
}
