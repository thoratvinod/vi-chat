package common

import (
	"fmt"
	"net/http"
)

type ErrorResponse interface {
	Error() string
	GetErrorCode() ErrorCode
}
type errorResponse struct {
	Code    ErrorCode
	Message string
}

type ErrorCode uint

const (
	DefaultErrorCode    = ErrInteralServerError
	DefaultErrorMessage = "Internal Error"
	DefaultStatusCode   = http.StatusInternalServerError
)

// Error codes
const (
	ErrInteralServerError ErrorCode = 1000 + iota
	ErrInvalidInput
	ErrLoginFailed
	ErrSignupFailed
	ErrDBOperationFailed
)

// Error messages by error codes
var errorMessages = map[ErrorCode]string{
	ErrInteralServerError: "Internal error",
	ErrInvalidInput:       "Received invalid input",
	ErrLoginFailed:        "Login failed",
	ErrDBOperationFailed:  "DB operation failed",
}

// HTTP status code by error codes
var httpMapper = map[ErrorCode]int{
	ErrInteralServerError: http.StatusInternalServerError,
}

func NewErrorResponse(code ErrorCode, message string) ErrorResponse {
	return &errorResponse{
		Code:    code,
		Message: message,
	}
}

func GetErrorMessage(errCode ErrorCode) string {
	msg, ok := errorMessages[errCode]
	if !ok {
		return DefaultErrorMessage
	}
	return msg
}

func GetStatusCode(errCode ErrorCode) int {
	statusCode, ok := httpMapper[errCode]
	if !ok {
		return DefaultStatusCode
	}
	return statusCode
}

func (err *errorResponse) Error() string {
	return fmt.Sprintf("%v;%v", GetErrorMessage(ErrorCode(err.Code)), err.Message)
}

func (err *errorResponse) GetErrorCode() ErrorCode {
	return err.Code
}
