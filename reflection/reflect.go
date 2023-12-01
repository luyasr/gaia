package reflection

import (
	"github.com/luyasr/gaia/log"
	"reflect"
)

func Must(obj any) (reflect.Type, reflect.Value) {
	valueOf := reflect.ValueOf(obj)
	if valueOf.Kind() != reflect.Ptr || valueOf.IsNil() {
		log.Fatal("must be a non null pointer to a struct")
	}

	valueOf = valueOf.Elem()
	typeOf := valueOf.Type()

	return typeOf, valueOf
}

func SetUp(obj any) {
	SetDefaultTag(obj)
}
