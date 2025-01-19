package validation

import (
	"encoding/json"
	"errors"

	rest_err "github.com/Moreno1337/api-ims/src/configuration/err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_BR_translation "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt_BR := pt_BR.New()
		unt := ut.New(pt_BR, pt_BR);
		transl, _ = unt.GetTranslator("pt_BR")
		pt_BR_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var validationErr validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Algum campo possui um tipo inválido")
	}

	if errors.As(validation_err, &validationErr) {
		errorCauses := []rest_err.Causes {}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Um ou mais campos inválidos", errorCauses)
	}

	err := "Erro ao tentar converter os campos, erro: " + validation_err.Error()
	return rest_err.NewBadRequestError(err)
}
