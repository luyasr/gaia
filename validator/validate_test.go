package validator

import (
	"testing"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Person struct {
	Name string `validate:"required" label:"姓名"`
	Age  int `validate:"required,isAdult" label:"年龄"`
}

func init() {
	Validate.RegisterValidation("isAdult", isAdult)
	Validate.RegisterTranslation("isAdult", Trans, registrationFunc, translateFunc)
}

func isAdult(fl validator.FieldLevel) bool {
	return fl.Field().Int() >= 18
}

var registrationFunc = func(ut ut.Translator) error {
    return ut.Add("isAdult", "{0}必须是成年人！", true) // see universal-translator for details
}

var translateFunc = func(ut ut.Translator, fe validator.FieldError) string {
    t, _ := ut.T("isAdult", fe.Field())
    return t
} 

func TestStruct(t *testing.T) {
	p := Person{
		// Name: "luya",
		Age:  14,
	}
	err := Struct(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
}
