package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type Header http.Header

type Endpoint struct {
	ID          uuid.UUID
	Name        string
	Description string
	Method      string
	URL         string
	Path        string
	Headers     Header
	Body        Body
	Response    Body
}

type Body map[string]any

func (b Body) IsEmpty() bool {
	return len(b) == 0
}

func (b Body) Byte() ([]byte, error) {
	return json.Marshal(b)
}

func (b Body) SetBody(body map[string]any) {
	b = body
}

func (b Body) Casting(endpointName string) string {
	typesCasting := NewTypesCasting(endpointName)
	return typesCasting.Cast(b)
}
