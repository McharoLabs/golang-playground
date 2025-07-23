package helpers

import (
	"reflect"
	"strings"
)

func GetJSONTag(model any, fieldName string) string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == fieldName {
			tag := field.Tag.Get("json")
			return strings.Split(tag, ",")[0]
		}
	}
	return fieldName
}

// ValidationErrors holds multiple field validation errors
type ValidationErrors map[string]string

func (v ValidationErrors) Error() string {
	var sb strings.Builder
	for field, msg := range v {
		sb.WriteString(field + ": " + msg + "; ")
	}
	return sb.String()
}
