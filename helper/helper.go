package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta
	Data interface{}
}

type Meta struct {
	Message string
	Code    int
	Status  string
}

func JsonResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
