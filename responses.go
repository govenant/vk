package vk

import "fmt"

type Error struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_msg"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("Error code %d: %s", err.Code, err.Message)
}

type ErrorResponse struct {
	Error *Error `json:"error"`
}

type AccessTokenResponse struct {
	ErrorResponse

	AccessToken string `json:"access_token"`
}
