package restapi

import (
	"bff/internal/infra/database/sqlc"
	"encoding/json"
	"fmt"
	"os"
)

const cachePath = "tmp/cache.json"

func SaveCache(endpoints []sqlc.Endpoint) error {
	dataCache := make(map[string]string, len(endpoints))
	for _, endpoint := range endpoints {
		dataCache[endpoint.Name] = fmt.Sprintf("%s%s", endpoint.Url, endpoint.Path)
	}

	cache, _ := json.Marshal(dataCache)
	if err := os.WriteFile(cachePath, cache, 0644); err != nil {
		return err
	}

	return nil
}

func GetUrlFromCache(urlPath string) (Url string, err error) {
	b, err := os.ReadFile(cachePath)
	if err != nil {
		return "", fmt.Errorf("error reading cache: %w", err)
	}

	endpoints := map[string]string{}
	if err := json.Unmarshal(b, &endpoints); err != nil {
		return "", fmt.Errorf("error unmarshal cache: %w", err)
	}

	url, ok := endpoints[urlPath]
	if !ok {
		return "", fmt.Errorf("url not found")
	}

	return url, nil
}
