package reflection

import (
	"reflect"

	"github.com/luyasr/gaia/errors"
)

func Must(obj any) (reflect.Type, reflect.Value, error) {
	valueOf := reflect.ValueOf(obj)
	if valueOf.Kind() != reflect.Ptr || valueOf.IsNil() {
		return nil, reflect.Value{}, errors.Internal("reflect", "must be a non null pointer to a struct")
	}

	valueOf = valueOf.Elem()
	typeOf := valueOf.Type()

	return typeOf, valueOf, nil
}

func SetUp(obj any) error {
	if err := SetDefaultTag(obj); err != nil {
		return err
	}

	return nil
}
