package tests

import (
	"bff/internal/management/dto/postman"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewEnvironment(t *testing.T) {
	tests := []struct {
		name           string
		filePath       string
		expectedEnv    postman.Environment
		domainExpected map[string]string
		expectedError  bool
	}{
		{
			name:     "Valid environment file",
			filePath: "/Users/isaacdsc/GolandProjects/bff/tmp/internal_FF.postman_environment.json",
			expectedEnv: postman.Environment{
				ID:   "7c9c1a8e-e25e-4fa2-9cf5-2f5953864034",
				Name: "Local",
				Values: []postman.Value{
					{Key: "host", Value: "http://localhost:3000", Type: "default", Enabled: true},
				},
				PostmanVariableScope: "environment",
				PostmanExportedAt:    time.Date(2024, 11, 30, 22, 28, 41, 939, time.UTC),
				PostmanExportedUsing: "Postman/11.18.1",
			},
			domainExpected: map[string]string{
				"host": "http://localhost:3000",
			},
			expectedError: false,
		},
		{
			name:           "Invalid environment file",
			filePath:       "temp/invalid_environment.json",
			expectedEnv:    postman.Environment{},
			expectedError:  true,
			domainExpected: make(map[string]string),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env, err := postman.NewEnvironment(tt.filePath)
			if (err != nil) != tt.expectedError {
				t.Errorf("NewEnvironment() error = %v, expectedError %v", err, tt.expectedError)
				return
			}

			assert.Equal(t, tt.domainExpected, env.ToDomain())
		})
	}
}
