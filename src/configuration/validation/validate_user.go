package valodation

import (
	"encoding/json"
	"errors"

	"github.com/DiogoLuizBonaparte/go_api.git/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ := unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErro validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationErro) {
		errosCauses := []rest_err.Causes{}

		for _, er := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: er.Translate(transl),
				Field:   er.Field(),
			}
			errosCauses = append(errosCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invelid", errosCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}

}
