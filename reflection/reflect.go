package reflection

import (
	"reflect"
)

func SetUp(obj any) error {
	if err := SetDefaultTag(obj); err != nil {
		return err
	}

	return nil
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
