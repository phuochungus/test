package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func createMessageByTag(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		switch err.Type().String() {
		case "string":
			return err.Field() + " is missing or empty"
		default:
			return err.Field() + "is missing or empty or zero"
		}
	default:
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			return err.Translate(translatorInstance)
		} else {
			return err.Error()
		}

	}
}

func ParseErrors(errors validator.ValidationErrors) []string {
	var msges []string
	for _, e := range errors {
		msges = append(msges, createMessageByTag(e))
	}
	return msges
}
