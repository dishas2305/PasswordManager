package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var trans ut.Translator

func init() {
	en := en.New()
	uni := ut.New(en, en)
	validate = validator.New()
	trans, _ = uni.GetTranslator("en")

	err := en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("Error in validate init", err.Error())
	}
}

func ValidateStruct(body interface{}) error {
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errMessages := validationErrors.Translate(trans)
		var msg string
		for _, err := range errMessages {
			if msg == "" {
				msg += err
			} else {
				msg += "\n" + err
			}
		}
		return errors.New(msg)
	}
	return nil
}
