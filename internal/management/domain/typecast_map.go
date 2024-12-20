package domain

import "reflect"

type Map struct{}

func (t Map) isMap(value any) bool {
	_, ok := value.(map[string]any)
	return ok
}

func (t Map) getStructureLine() string {
	return "%s []%s `json:\"%s\"` \n"
}

func (t Map) getTypeMapped(value map[string]any) string {
	var strType string
	for _, v := range value {
		strType = reflect.TypeOf(v).String()
	}
	return strType
}
