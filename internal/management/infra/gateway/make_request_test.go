package gateway

import (
	"bff/internal/management/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestApi_Request(t *testing.T) {
	tests := []struct {
		name       string
		item       domain.Endpoint
		mockServer func() *httptest.Server
		wantErr    bool
	}{
		{
			name: "successful request",
			item: domain.Endpoint{
				Method:  http.MethodGet,
				URL:     "/test",
				Headers: map[string][]string{"Content-Type": {"application/json"}},
			},
			mockServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, http.MethodGet, r.Method)
					assert.Equal(t, "/test", r.URL.Path)
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"key":"value"}`))
				}))
			},
			wantErr: false,
		},
		{
			name: "request with error",
			item: domain.Endpoint{
				Method:  http.MethodGet,
				URL:     "/test",
				Headers: map[string][]string{"Content-Type": {"application/json"}},
			},
			mockServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := tt.mockServer()
			defer server.Close()

			tt.item.URL = server.URL + tt.item.URL

			api := NewRequestApi()
			_, err := api.Request(context.Background(), tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestApi.Request() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
