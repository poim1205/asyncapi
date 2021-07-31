package asyncapi2

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("version", validateVersionNumber)
}

func ValidateStruct(i interface{}) error {
	err := validate.Struct(i)
	var errMessage string
	if err != nil {
		// TODO: define how you want return the error, below an example
		for _, err := range err.(validator.ValidationErrors) {
			errMessage = fmt.Sprintf("%v should be %v %v, current value is %v ",
				err.Namespace(),
				err.ActualTag(),
				err.Param(),
				err.Value(),
			)
		}
		return errors.New(errMessage)
	}
	return nil
}

func validateVersionNumber(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	if len(name) == 0 {
		return false
	}

	regex, _ := regexp.Compile("^(\\d+\\.)?(\\d+\\.)?(\\d+)$")
	result := regex.MatchString(name)
	return result
}
