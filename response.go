package responsepkg

import "net/http"

type Resp interface{
	Status() int
	Message() string
	Data() any
	Response() string
}

type response struct{
	RespStatus int `json:"status"`
	RespMessage string `json:"message"`
	RespData any `json:"data"`
	RespResponse string `json:"response"`
}

func (r response) Status() int {
	return r.RespStatus
}

func (r response) Message() string {
	return r.RespMessage
}

func (r response) Data() any {
	return r.RespData
}

func (r response) Response() string {
	return r.RespResponse
}

func ResponseCreated(message string, data any) Resp {
	result := response{
		RespStatus: http.StatusCreated,
		RespMessage: message,
		RespData: data,
		RespResponse: "Created Successfully",
	}

	return result
}

func ResponseSuccess(message string, data any) Resp {
	result := response{
		RespStatus: http.StatusOK,
		RespMessage: message,
		RespData: data,
		RespResponse: "Success",
	}

	return result
}

func ResponseNoData(message string) Resp {
	result := response{
		RespStatus: http.StatusNoContent,
		RespMessage: message,
		RespData: nil,
		RespResponse: "No data Found",
	}

	return result
}