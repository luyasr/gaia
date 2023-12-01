package validator

import (
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/luyasr/gaia/log"
	"reflect"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = validator.New()
	// english translator
	enTranslator := en.New()
	// chinese translator
	zhTranslator := zh.New()

	uni := ut.New(enTranslator, zhTranslator)

	// get the language you need
	trans, _ = uni.GetTranslator("zhTranslator")

	// Register a function to get the custom label in the struct tag as the field name
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// the authenticator registers the translator
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatalf("authenticator registration translator error: %v", err)
	}
}

func Struct(target any) error {
	err := validate.Struct(target)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return errors.New(err.Error())
		}

		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(e.Translate(trans))
		}
	}
	return nil
}
