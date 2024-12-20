package postman

import (
	"bff/internal/management/util"
	"fmt"
	"time"
)

type Environment struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Values               []Value   `json:"values"`
	PostmanVariableScope string    `json:"_postman_variable_scope"`
	PostmanExportedAt    time.Time `json:"_postman_exported_at"`
	PostmanExportedUsing string    `json:"_postman_exported_using"`
}

func NewEnvironment(filePath string) (Environment, error) {
	var env Environment
	if err := util.ReadFile(filePath, &env); err != nil {
		return Environment{}, fmt.Errorf("error converter to JSON: %v", err)
	}

	return env, nil
}

func (e Environment) ToDomain() map[string]string {
	var envVariablesMap = make(map[string]string)
	for _, value := range e.Values {
		envVariablesMap[value.Key] = value.Value
	}

	return envVariablesMap
}

type Value struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

type ByteEnvironment []byte
