package helper

import "golang.org/x/crypto/bcrypt"

type Desired_Output struct {
	Message string `json: "message"`
	Code    int    `json: "code"`
}

type Response struct {
	Code    int         `json: "code"`
	Message string      `json: "message"`
	Data    interface{} `json: "data"`
}

func ResponsesAuth(message string, code int, data interface{}) Response {
	jsonResponse := Response{
		Code: code,
		Data: data,
	}
	return jsonResponse
}

func ResponsesFormat(message string, code int, data interface{}) Response {
	jsonResponse := Response{
		Code:    code,
		Message: message,
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
