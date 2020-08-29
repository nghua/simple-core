package validate

import (
	"errors"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	log "github.com/sirupsen/logrus"
)

var validate *validator.Validate
var translator ut.Translator

func init() {
	validate = validator.New()
	uni := ut.New(zh_Hans_CN.New())
	translator, _ = uni.GetTranslator("zh_Hans_CN")

	err := zh.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		log.Printf("验证器翻译出错：%v", err)
	}
}

func Struct(s interface{}) error {
	validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		label := f.Tag.Get("label")
		return label
	})

	err := validate.Struct(s)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return errors.New(v.Translate(translator))
		}
	}

	return nil
}
