package domain

import (
	"fmt"
	"reflect"
)

type Slice struct{}

func (t Slice) isSlice(value any) bool {
	_, ok := value.([]any)
	return ok
}

func (t Slice) getTypeSliced(value []any) string {
	var strType string
	for i, v := range value {
		strType = reflect.TypeOf(v).String()
		if i > 0 {
			before := value[i-1]
			// identify type of []any
			if reflect.TypeOf(before) != reflect.TypeOf(v) {
				fmt.Println("compare:::", reflect.TypeOf(before), reflect.TypeOf(v))
				return "any"
			}
		}
	}

	return strType
}

func (t Slice) getStructureLine() string {
	return "%s []%s `json:\"%s\"` \n"
}
