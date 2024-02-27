package reflection

import (
	"reflect"
)

type Config struct {
	IgnoreEmpty bool
}

// StructToMap converts a struct to a map using reflection.
// If you want to remove default zero values, you can use pointer types in the struct.
func StructToMap(obj any) map[string]any {
	valueOf := reflect.ValueOf(obj)

	// If the object is a pointer, dereference it.
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	typeOf := valueOf.Type()
	typeOfNumField := typeOf.NumField()
	resultMap := make(map[string]any, typeOfNumField)

	// Iterate over all fields of the struct.
	for i := 0; i < typeOfNumField; i++ {
		tField, vField := typeOf.Field(i), valueOf.Field(i)

		// If the field is a pointer, dereference it.
		if vField.Kind() == reflect.Ptr {
			if vField.IsNil() {
				continue
			}
			vField = vField.Elem()
		}

		fieldValue := vField.Interface()
		fieldName := tField.Name
		if tag, ok := tField.Tag.Lookup("json"); ok {
			fieldName = tag
		}

		switch vField.Kind() {
		case reflect.Struct:
			resultMap[fieldName] = StructToMap(fieldValue)
		default:
			resultMap[fieldName] = fieldValue
		}
	}

	return resultMap
}
