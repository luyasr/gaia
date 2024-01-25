package reflection

import "reflect"

func StructToMap(obj any) map[string]any {
	return structToMapInternal(obj, false)
}

func StructToMapRmEmpty(obj any) map[string]any {
	return structToMapInternal(obj, true)
}

func structToMapInternal(obj any, removeEmpty bool) map[string]any {
	valueOf := reflect.ValueOf(obj)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}
	typeOf := valueOf.Type()

	m := make(map[string]any)

	for i := 0; i < typeOf.NumField(); i++ {
		tField, vField := typeOf.Field(i), valueOf.Field(i)
		fieldValue := vField.Interface()

		if removeEmpty && fieldValue == reflect.Zero(vField.Type()).Interface() {
			continue
		}

		if tag, ok := tField.Tag.Lookup("json"); ok {
			m[tag] = fieldValue
		} else {
			m[tField.Name] = fieldValue
		}
	}

	return m
}