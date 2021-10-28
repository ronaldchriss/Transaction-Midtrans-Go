package helper

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
