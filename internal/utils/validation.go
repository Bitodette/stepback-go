package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ParseBody(r *http.Request, dst interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return false
	}
	return true
}

// runs all `validate:"..."` struct tags
// returns map of field name to error msg, nil if all good
// NOTE: field names are lowercase
func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	result := make(map[string]string)
	for _, e := range validationErrs {
		field := strings.ToLower(e.Field())
		result[field] = parseTag(e.Tag(), e.Param())
	}
	return result
}

func parseTag(tag, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Must be a valid email"
	case "min":
		return "Must be at least " + param + " characters"
	case "max":
		return "Must be at most " + param + " characters"
	case "len":
		return "Must be exactly " + param + " characters"
	case "gte":
		return "Must be greater than or equal to " + param
	case "lte":
		return "Must be less than or equal to " + param
	default:
		return "Invalid value"
	}
}
