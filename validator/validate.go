package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/luyasr/gaia/log"
)

var (
	Validate *validator.Validate
	TransEn  ut.Translator
	TransZh  ut.Translator
)

func init() {
	initValidator()
}

// initValidator initializes the validator and translator.
func initValidator() {
	Validate = validator.New()

	// Set up English and Chinese translators.
	enTranslator := en.New()
	zhTranslator := zh.New()

	uni := ut.New(enTranslator, zhTranslator)

	// Get the English and Chinese translators.
	TransEn, _ = uni.GetTranslator("en")
	TransZh, _ = uni.GetTranslator("zh")

	// Register the default translations for the English translator.
	if err := enTrans.RegisterDefaultTranslations(Validate, TransEn); err != nil {
		log.Fatalf("authenticator registration translator error: %v", err)
	}

	// Register the default translations for the Chinese translator.
	if err := zhTrans.RegisterDefaultTranslations(Validate, TransZh); err != nil {
		log.Fatalf("authenticator registration translator error: %v", err)
	}

	// Register the tag name function.
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			name = ""
		}
		label := fld.Tag.Get("label")
		if label != "" {
			name = label
		}

		return name
	})
}

func RegisterValidation(tag string, fn validator.Func) error {
	return Validate.RegisterValidation(tag, fn)
}

func RegisterTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc, translationFn validator.TranslationFunc) error {
	return Validate.RegisterTranslation(tag, trans, registerFn, translationFn)
}

// Struct validates the given struct using the validator and translator.
func Struct(target any) error {
	return StructWithLang(target, "zh")
}

// StructWithLang validates the given struct using the validator and translator with a specific language.
func StructWithLang(target any, language string) error {
	err := Validate.Struct(target)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return errors.New(err.Error())
		}

		var trans ut.Translator
		switch language {
		case "en":
			trans = TransEn
		case "zh":
			trans = TransZh
		default:
			trans = TransEn
		}

		// Translate the validation errors.
		sb := strings.Builder{}
		for _, e := range err.(validator.ValidationErrors) {
			if sb.Len() > 0 {
				sb.WriteString("; ")
			}
			sb.WriteString(e.Translate(trans))
		}
		return errors.New(sb.String())
	}
	return nil
}
