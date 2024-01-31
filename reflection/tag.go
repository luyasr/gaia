package reflection

import (
	"reflect"
	"strconv"

	"github.com/luyasr/gaia/errors"
)

func must(obj any) (reflect.Type, reflect.Value, error) {
	valueOf := reflect.ValueOf(obj)
	if valueOf.Kind() != reflect.Ptr || valueOf.IsNil() {
		return nil, reflect.Value{}, errors.Internal("reflect", "must be a non null pointer to a struct")
	}

	valueOf = valueOf.Elem()
	typeOf := valueOf.Type()

	return typeOf, valueOf, nil
}

func SetDefaultTag(obj any) error {
	typeOf, valueOf, err := must(obj)
	if err != nil {
		return err
	}

	for i := 0; i < typeOf.NumField(); i++ {
		tFiled, vFiled := typeOf.Field(i), valueOf.Field(i)
		tag, ok := tFiled.Tag.Lookup("default")
		if vFiled.IsZero() && ok {
			switch vFiled.Kind() {
			case reflect.String:
				vFiled.SetString(tag)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				parseInt, err := strconv.ParseInt(tag, 10, 64)
				if err != nil {
					return errors.Internal("reflection setting default failed", "error parsing int: %s", err)
				}
				vFiled.SetInt(parseInt)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				parseUint, err := strconv.ParseUint(tag, 10, 64)
				if err != nil {
					return errors.Internal("reflection setting default failed", "error parsing uint: %s", err)
				}
				vFiled.SetUint(parseUint)
			case reflect.Float32, reflect.Float64:
				parseFloat, err := strconv.ParseFloat(tag, 64)
				if err != nil {
					return errors.Internal("reflection setting default failed", "error parsing float: %s", err)
				}
				vFiled.SetFloat(parseFloat)
			case reflect.Bool:
				parseBool, err := strconv.ParseBool(tag)
				if err != nil {
					return errors.Internal("reflection setting default failed", "error parsing bool: %s", err)
				}
				vFiled.SetBool(parseBool)
			case reflect.Ptr:
				vFiled.Set(reflect.New(vFiled.Type().Elem()))
			case reflect.Struct:
				SetDefaultTag(vFiled.Addr().Interface())
			default:
				return errors.Internal("reflection setting default failed", "unsupported type: %s", vFiled.Kind())
			}
		}
	}

	return nil
}
