package gateway

import (
	"bff/internal/management/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestApi struct {
	client *http.Client
}

func NewRequestApi() *RequestApi {
	return &RequestApi{
		client: &http.Client{},
	}
}

func (g RequestApi) Request(ctx context.Context, item domain.Endpoint) (map[string]any, error) {
	var output map[string]any

	var (
		req *http.Request
		err error
	)

	if !item.Body.IsEmpty() {
		p, err := item.Body.Byte()
		if err != nil {
			return output, fmt.Errorf("error on marshal payload: %w", err)
		}
		req, err = http.NewRequest(item.Method, item.URL, bytes.NewReader(p))
		if err != nil {
			return output, fmt.Errorf("error on create new request: %w", err)
		}
	} else {
		req, err = http.NewRequest(item.Method, item.URL, nil)
		if err != nil {
			return output, fmt.Errorf("error on create new request: %w", err)
		}
	}

	for key, values := range item.Headers {
		for i := range values {
			req.Header.Add(key, values[i])
		}
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return output, fmt.Errorf("error on call httpclient: %w", err)
	}

	defer resp.Body.Close()

	// Ler a resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return output, fmt.Errorf("error read response httpclient: %w", err)
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return output, fmt.Errorf("error unmarshal response httpclient: %w", err)
	}

	return output, nil
}
