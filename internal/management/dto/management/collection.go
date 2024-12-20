package management

import (
	"bff/internal/management/domain"
	"bff/internal/management/util"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type CollectionDto struct {
	Info Info          `json:"info"`
	Item []PostmanItem `json:"item"`
}

func NewPostmanCollection(filePath string) (CollectionDto, error) {
	var pc CollectionDto
	if err := util.ReadFile(filePath, &pc); err != nil {
		return CollectionDto{}, fmt.Errorf("error converter to JSON: %v", err)
	}

	return pc, nil
}

func (pc CollectionDto) ToDomain(envVariables map[string]string) ([]domain.Endpoint, error) {
	output := make([]domain.Endpoint, len(pc.Item))
	for i, item := range pc.Item {
		item.Request.URL.Raw = item.Request.URL.Raw.replaceToValue(envVariables)
		r, err := item.toDomain(envVariables)
		if err != nil {
			return nil, fmt.Errorf("error on convert to domain: %v", err)
		}
		output[i] = r
	}

	return output, nil
}

type Info struct {
	PostmanID  string `json:"_postman_id"`
	Name       string `json:"name"`
	Schema     string `json:"schema"`
	ExporterID string `json:"_exporter_id"`
}

type PostmanItem struct {
	Name     string         `json:"name"`
	Request  PostmanRequest `json:"request"`
	Response []interface{}  `json:"response" gorm:"-"`
}

func (p PostmanItem) toDomain(envVariables map[string]string) (domain.Endpoint, error) {
	path := ""
	for _, pathName := range p.Request.URL.Path {
		path += fmt.Sprintf("/%s", pathName)
	}

	var result map[string]any
	if p.Request.Body.Raw != "" {
		jsonBytes := []byte(p.Request.Body.Raw)
		if err := json.Unmarshal(jsonBytes, &result); err != nil {
			return domain.Endpoint{}, fmt.Errorf("error on unmarshal body: %v", err)
		}
	}

	return domain.Endpoint{
		ID:          uuid.New(),
		Name:        p.Name,
		Description: "",
		Method:      p.Request.Method,
		URL:         p.Request.URL.Raw.String(),
		Path:        path,
		Headers:     p.Request.getDomainHeader(),
		Body:        result,
	}, nil
}

func (p PostmanItem) replaceValues(envVariables map[string]string) {
	for _, header := range p.Request.Header {
		header.replaceToValue(envVariables)
	}

	p.Request.URL.Raw = p.Request.URL.Raw.replaceToValue(envVariables)
}

type PostmanRequest struct {
	Method string             `json:"method"`
	Header []Header           `json:"header" gorm:"foreignKey:PostmanItemID"`
	URL    URL                `json:"url"`
	Body   PostmanRequestBody `json:"body"`
}

type PostmanRequestBody struct {
	Mode    string `json:"mode"`
	Raw     string `json:"raw"`
	Options struct {
		Raw struct {
			Language string `json:"language"`
		} `json:"raw"`
	} `json:"options"`
}

func (pr PostmanRequest) getDomainHeader() (output domain.Header) {
	for _, header := range pr.Header {
		output = header.ToDomain()
	}

	return
}

type Header struct {
	PostmanItemID uint
	Key           string   `json:"key"`
	Value         KeyValue `json:"value"`
	Type          string   `json:"type"`
}

func (h Header) ToDomain() domain.Header {
	return map[string][]string{h.Key: {h.Value.String()}}
}

func (h Header) replaceToValue(envVariables map[string]string) {
	h.Value = h.Value.replaceToValue(envVariables)
}

type KeyValue string

func (r KeyValue) String() string {
	return string(r)
}

func (r KeyValue) replaceToValue(envVariables map[string]string) KeyValue {
	for key, val := range envVariables {
		placeholder := fmt.Sprintf("{{%s}}", key)
		return KeyValue(strings.ReplaceAll(string(r), placeholder, val))
	}

	return r
}

type URL struct {
	Raw  KeyValue `json:"raw"`
	Host []string `json:"host"`
	Path []string `json:"path"`
}

type ByteCollection []byte
