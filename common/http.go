package common

import (
	"demo/model"
)

func HttpResponse(code int, data any, message string) model.HttpResponse {
	return model.HttpResponse{
		StatusCode: code,
		Msg:        message,
		Data:       data,
	}
}
