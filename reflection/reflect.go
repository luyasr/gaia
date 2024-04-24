package reflection

import (
	"reflect"
)

func SetUp(obj any) error {
	return SetDefaultTag(obj)
}

func GetFieldValue(structure any, fieldName string) (any, bool) {
	v := reflect.ValueOf(structure)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, false
	}
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, false
	}
	return field.Interface(), true
}

func GetFieldValueByType(structure any, fieldType reflect.Type) map[string]any {
	v := reflect.ValueOf(structure)
	fields := make(map[string]any)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fields
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Type() == fieldType {
			fields[v.Type().Field(i).Name] = field.Interface()
		}
	}

	return fields
}
