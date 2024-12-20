package tests

import (
	"fmt"
	"net/http"
	"testing"
)

var verbs = map[string]string{
	"GET":    "query",
	"POST":   "mutation",
	"PUT":    "mutation",
	"DELETE": "mutation",
	"PATCH":  "mutation",
}

func TestJsonToStruct(t *testing.T) {
	//method := http.MethodGet
	method := http.MethodPost

	payload := map[string]any{
		"getInternalFF": map[string]any{
			"_postman_id": "f4b3b2d1-2c7d-4e7b-9b4d-445f9c9b2e2b",
		},
	}

	gqlQuery := convertToGraphQL(payload, method)
	fmt.Println(gqlQuery)
}

func convertToGraphQL(payload map[string]any, method string) string {

	gql := fmt.Sprintf("%s {", verbs[method])

	for key, value := range payload {
		gql += fmt.Sprintf(" %s {", key)
		for subKey, subValue := range value.(map[string]any) {
			gql += fmt.Sprintf(" %s: \"%v\"", subKey, subValue)
		}
		gql += " }"
	}
	gql += " }"
	return gql
}
