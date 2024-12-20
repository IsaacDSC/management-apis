package restapi

import (
	"bff/internal/management/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func ProxyHttpClient(ctx context.Context, item domain.Endpoint) (map[string]any, error) {
	output := map[string]any{}
	client := new(http.Client)

	var (
		req *http.Request
		err error
	)
	if item.Body.IsEmpty() {
		req, err = http.NewRequest(item.Method, item.URL, nil)
		if err != nil {
			return output, fmt.Errorf("error on create new request: %w", err)
		}
	} else {
		p, err := item.Body.Byte()
		if err != nil {
			return output, fmt.Errorf("error on marshal payload: %w", err)
		}
		req, err = http.NewRequest(item.Method, item.URL, bytes.NewReader(p))
		if err != nil {
			return output, fmt.Errorf("error on create new request: %w", err)
		}
	}

	req.Header.Add("user-agent", "proxy-bff")
	for key, values := range item.Headers {
		for i := range values {
			req.Header.Add(key, values[i])
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return output, fmt.Errorf("error on call httpclient: %w", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return output, fmt.Errorf("error on decode response: %w", err)
	}

	return output, err
}
