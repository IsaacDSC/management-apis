package domain

import (
	"fmt"
)

type Gql struct{}

func (g *Gql) StructToGqlType(api API) string {
	mapper := map[string]string{
		"GET":    "Query",
		"POST":   "Mutation",
		"PUT":    "Mutation",
		"DELETE": "Mutation",
		"PATCH":  "Mutation",
	}

	//typeMapper := map[any]string{
	//	uuid.UUID{}: "ID",
	//	string:      "String",
	//	int:         "Int",
	//	int64:       "Int",
	//	float64:     "Float",
	//}

	//TODO: tipar dados de resposta

	typeTxt := ""
	for i := range api.Endpoints {
		typeTxt += fmt.Sprintf("type %s {\n", api.Endpoints[i].Name)
		typeTxt += fmt.Sprintf("%s: ")
		typeTxt += fmt.Sprintf("}\n")
	}

	//TODO: tipar dados de input

	//TODO: tipar query and mutations
	typeTxt += "\n\n"
	for i := range api.Endpoints {
		typeTxt += fmt.Sprintf("type %s {\n", mapper[api.Endpoints[i].Method])
		typeTxt += fmt.Sprintf("%s(%s!): %s!", api.Endpoints[i].Name, "", api.Endpoints[i].Body)
		typeTxt += fmt.Sprintf("}\n")
	}

	return ""
}
