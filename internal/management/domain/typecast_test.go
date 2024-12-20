package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestTypeCasting(t *testing.T) {
	tests := []struct {
		name     string
		response map[string]any
		expected string
	}{
		{
			name: "response with various types",
			response: map[string]any{
				"stringKey": "stringValue",
				"intKey":    123,
				"floatKey":  123.45,
				"boolKey":   true,
				"sliceKey":  []any{"item1", "item2"},
				"sliceKey2": []any{"item1", 123},
				"complex":   map[string]any{"complex_key": map[string]any{"complex_keyValue": "complexValue"}},
				"mapKey":    map[string]any{"nestedKey": "nestedValue"},
			},
			expected: "type Response struct {\nSliceKey2 []any `json:\"sliceKey2\"` \ntype MapKey struct {\ntype Response struct {\nSliceKey2 []any `json:\"sliceKey2\"` \ntype MapKey struct {\nNestedKey string `json:\"nestedKey\"` \n}\n} `json:\"mapKey\"` \nStringKey string `json:\"stringKey\"` \nIntKey int `json:\"intKey\"` \nFloatKey float64 `json:\"floatKey\"` \nBoolKey bool `json:\"boolKey\"` \nSliceKey []string `json:\"sliceKey\"` \n}",
		},
		{
			name: "response with unknown type",
			response: map[string]any{
				"unknownKey": struct{}{},
			},
			expected: "type,",
		},
	}

	defer func() {
		os.Remove("./tmp/main.go")
		os.Remove("./tmp/main")
	}()

	funcMain := ` func main(){println("ok")}`
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewTypesCasting("teste")
			cast := e.Cast(tt.response)
			cast = fmt.Sprintf("package main\n\n%s\n\n%s", funcMain, cast)
			assert.NoError(t, os.WriteFile("./tmp/main.go", []byte(cast), 0644))
			cmd := exec.Command("go", "build", "-o", "./tmp/main", "./tmp/main.go")
			assert.NoError(t, cmd.Run())
		})
	}
}
