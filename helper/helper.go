package helper

type Desired_Output struct {
	Message string `json: "message"`
	Code    int    `json: "code"`
}

type Response struct {
	DO   Desired_Output `json: "desired_output"`
	Data interface{}    `json: "data"`
}

func ResponsesFormat(message string, code int, data interface{}) Response {
	do := Desired_Output{
		Message: message,
		Code:    code,
	}
	jsonResponse := Response{
		DO:   do,
		Data: data,
	}
	return jsonResponse
}
