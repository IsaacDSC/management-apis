package service

import (
	"bff/internal/management/domain"
	"bff/mocks"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_RegistryApi(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repository := mocks.NewMockManagementAdapterRepository(controller)
	gateway := mocks.NewMockRequestApiAdapter(controller)

	svc := NewService(gateway, repository)

	tests := []struct {
		name          string
		api           domain.API
		setupMock     func()
		expectedError string
	}{
		{
			name: "successful registry",
			setupMock: func() {
				repository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
				gateway.EXPECT().Request(gomock.Any(), gomock.Any()).Return(map[string]any{}, nil)
			},
			api: domain.API{
				Endpoints: []domain.Endpoint{
					{Name: "testEndpoint", Response: domain.Body{}},
				},
			},
			expectedError: "",
		},
		{
			name: "request error",
			setupMock: func() {
				gateway.EXPECT().Request(gomock.Any(), gomock.Any()).Return(map[string]any{}, errors.New("request error"))
			},
			api: domain.API{
				Endpoints: []domain.Endpoint{
					{Name: "testEndpoint", Body: domain.Body{}},
				},
			},
			expectedError: "error making request: request error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			err := svc.RegistryApi(context.Background(), tt.api)
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}

		})
	}
}
