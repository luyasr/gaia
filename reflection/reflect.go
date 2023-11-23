package reflection

import (
	"log"
	"reflect"
)

func Must(obj any) (reflect.Type, reflect.Value) {
	valueOf := reflect.ValueOf(obj)
	if valueOf.Kind() != reflect.Ptr || valueOf.IsNil() {
		log.Fatal("must be a non-nil pointer to a struct")
		return nil, reflect.Value{}
	}

	valueOf = valueOf.Elem()
	typeOf := valueOf.Type()

	return typeOf, valueOf
}

func SetUp(obj any) {
	SetDefaultTag(obj)
}
