package pkg

import (
	"context"
	"github.com/google/uuid"
	"testing"
)

func TestSetAndGetFromCtx(t *testing.T) {
	tests := []struct {
		name        string
		key         any
		correlation uuid.UUID
	}{
		{
			name:        "Test with valid UUID",
			key:         "correlationID",
			correlation: uuid.New(),
		},
		{
			name:        "Test with another valid UUID",
			key:         "anotherCorrelationID",
			correlation: uuid.New(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctx = SetToCtx(ctx, tt.key, tt.correlation)
			got := GetFromCtx(ctx, tt.key)
			if got != tt.correlation {
				t.Errorf("GetFromCtx() = %v, want %v", got, tt.correlation)
			}
		})
	}
}
