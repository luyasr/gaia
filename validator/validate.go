package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/luyasr/gaia/log"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
)

func init() {
	setupValidator()
}

// setupValidator initializes the validator and translator.
func setupValidator() {
	Validate = validator.New()

	// Set up English and Chinese translators.
	enTranslator := en.New()
	zhTranslator := zh.New()

	uni := ut.New(enTranslator, zhTranslator)

	// Get the Chinese translator.
	Trans, _ = uni.GetTranslator("zhTranslator")

	// Register a function to get the custom label in the struct tag as the field name.
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// Register the Chinese translations for the validator.
	err := zhTrans.RegisterDefaultTranslations(Validate, Trans)
	if err != nil {
		log.Fatalf("authenticator registration translator error: %v", err)
	}
}

func RegisterValidation(tag string, fn validator.Func) error {
	return Validate.RegisterValidation(tag, fn)
}

func RegisterTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc, translationFn validator.TranslationFunc) error {
	return Validate.RegisterTranslation(tag, trans, registerFn, translationFn)
}

// Struct validates the given struct using the validator and translator.
func Struct(target any) error {
	err := Validate.Struct(target)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return errors.New(err.Error())
		}

		// Translate the validation errors.
		var errs []string
		for _, e := range err.(validator.ValidationErrors) {
			errs = append(errs, e.Translate(Trans))
		}
		return errors.New(strings.Join(errs, "; "))
	}
	return nil
}
