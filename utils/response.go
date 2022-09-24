package utils

type Response struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(data interface{}, message interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}
