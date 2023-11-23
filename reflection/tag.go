package reflection

import (
	"reflect"
	"strconv"
)

func SetDefaultTag(obj any) {
	typeOf, valueOf := Must(obj)

	for i := 0; i < typeOf.NumField(); i++ {
		tFiled, vFiled := typeOf.Field(i), valueOf.Field(i)
		tag, _ := tFiled.Tag.Lookup("default")
		if vFiled.IsZero() {
			switch vFiled.Kind() {
			case reflect.String:
				vFiled.SetString(tag)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				parseInt, _ := strconv.ParseInt(tag, 10, 64)
				vFiled.SetInt(parseInt)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				parseUint, _ := strconv.ParseUint(tag, 10, 64)
				vFiled.SetUint(parseUint)
			case reflect.Float32, reflect.Float64:
				parseFloat, _ := strconv.ParseFloat(tag, 64)
				vFiled.SetFloat(parseFloat)
			case reflect.Bool:
				parseBool, _ := strconv.ParseBool(tag)
				vFiled.SetBool(parseBool)
			case reflect.Struct:
				SetDefaultTag(vFiled.Addr().Interface())
			}
		}
	}
}
