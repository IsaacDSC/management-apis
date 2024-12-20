package domain

import (
	"fmt"
	"strings"
)

type TypesCasting struct {
	castStr string
	Slice
	Map
}

func NewTypesCasting(endpointName string) *TypesCasting {
	return &TypesCasting{
		castStr: fmt.Sprintf("type %s struct {\n", strings.ToUpper(endpointName[:1])+endpointName[1:]),
	}
}

func (t TypesCasting) Cast(input map[string]any) string {
	for key, value := range input {
		Key := strings.ToUpper(key[:1]) + key[1:]
		if t.isPrimitive(value) {
			t.castStr += fmt.Sprintf("%s %T `json:\"%s\"` \n", Key, value, key)
		} else {
			if t.isMap(value) {
				t.castStr += fmt.Sprintf("%s struct {\n", Key)
				t.castStr += t.T(value.(map[string]any))
				t.castStr += fmt.Sprintf("} `json:\"%s\"` \n", key)
			}
			if t.isSlice(value) {
				t.castStr += fmt.Sprintf(t.Slice.getStructureLine(), Key, t.getTypeSliced(value.([]any)), key)
			}
		}
	}
	t.castStr += "}"
	return t.castStr
}

func (t TypesCasting) T(input map[string]any) string {
	var text string
	for key, value := range input {
		Key := strings.ToUpper(key[:1]) + key[1:]
		if t.isPrimitive(value) {
			text += fmt.Sprintf("%s %T `json:\"%s\"` \n", Key, value, key)
		} else {
			if t.isMap(value) {
				text += t.T(value.(map[string]any))
			}
			if t.isSlice(value) {
				text += fmt.Sprintf(t.Slice.getStructureLine(), Key, t.getTypeSliced(value.([]any)), key)
			}
		}
	}
	return text
}

func (t TypesCasting) isPrimitive(value any) bool {
	switch value.(type) {
	case string, int, float64, bool:
		return true
	}
	return false
}
