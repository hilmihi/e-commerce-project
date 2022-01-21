package helper

import "golang.org/x/crypto/bcrypt"

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
