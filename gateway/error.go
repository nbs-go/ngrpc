package gateway

type ErrorData struct {
	Debug *ErrorDebugData `json:"debug,omitempty"`
}

type ErrorDebugData struct {
	Message  string            `json:"message"`
	Traces   []string          `json:"traces"`
	Metadata map[string]string `json:"metadata"`
}

var InternalError = Response{
	Success: false,
	Code:    "500",
	Message: "Internal Error",
	Data:    nil,
}

var NotFoundError = Response{
	Success: false,
	Code:    "404",
	Message: "Not Found",
	Data:    nil,
}

var BadRequestError = Response{
	Success: false,
	Code:    "400",
	Message: "Bad Request",
	Data:    nil,
}

const internalErrorJson = `{"success": false, "code": "500", "message": "Internal Error", "data": null}`
