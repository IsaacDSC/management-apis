package restapi

import (
	"bff/internal/management/domain"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func DefaultProxyHandler(w http.ResponseWriter, r *http.Request) error {
	name := r.PathValue("endpoint_name")

	defer r.Body.Close()

	var body map[string]any
	if r.Method != http.MethodGet {
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			//TODO: mover para retornar um json
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("invalid body"))
			return err
		}
	}

	url, err := GetUrlFromCache(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("url not found"))
		return err
	}

	resp, err := ProxyHttpClient(r.Context(), domain.Endpoint{
		ID:      uuid.New(),
		Name:    name,
		Method:  r.Method,
		URL:     url,
		Headers: domain.Header(r.Header),
		Body:    body,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("error on call httpclient"))
		return err
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("error on encode response"))
		return err
	}

	return json.NewEncoder(w).Encode(resp)
}
