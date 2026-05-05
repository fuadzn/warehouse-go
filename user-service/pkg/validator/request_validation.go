package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(data interface{}) error {
	var errorMessages []string

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errorMessages = append(errorMessages, err.Field()+" is required")
			case "email":
				errorMessages = append(errorMessages, err.Field()+" is not valid")
			case "min":
				errorMessages = append(errorMessages, err.Field()+" must be at least "+err.Param()+" characters")
			case "max":
				errorMessages = append(errorMessages, err.Field()+" must be at most "+err.Param()+" characters")
			default:
				errorMessages = append(errorMessages, err.Error())
			}
		}
		return errors.New("failed validation: " + joinMessage(errorMessages))
	}

	return nil
}

func joinMessage(errorMessages []string) string {
	result := ""
	for i, message := range errorMessages {
		if i > 0 {
			result += ", "
		}
		result += message
	}
	return result
}
