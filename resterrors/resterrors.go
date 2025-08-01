package resterrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestError interface {
	Message() string
	DispMessage() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage     string        `json:"message"`
	DisplayMessage string        `json:"display_message"`
	ErrStatus      int           `json:"status"`
	ErrError       string        `json:"error"`
	ErrCauses      []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) DispMessage() string {
	return e.DisplayMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

func NewRestError(message string, displayMessage string, status int, err string, causes []interface{}) RestError {
	return &restErr{
		ErrMessage:     message,
		DisplayMessage: displayMessage,
		ErrStatus:      status,
		ErrError:       err,
		ErrCauses:      causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestError, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string, displayMessage string) RestError {
	return restErr{
		ErrMessage:     message,
		DisplayMessage: displayMessage,
		ErrStatus:      http.StatusBadRequest,
		ErrError:       "bad_request",
	}
}

func NewNotFoundError(message string, displayMessage string) RestError {
	return restErr{
		ErrMessage:     message,
		DisplayMessage: displayMessage,
		ErrStatus:      http.StatusNotFound,
		ErrError:       "not_found",
	}
}

func NewUnauthorizedError(message string, displayMessage string) RestError {
	return restErr{
		ErrMessage:     message,
		DisplayMessage: displayMessage,
		ErrStatus:      http.StatusUnauthorized,
		ErrError:       "unauthorized",
	}
}

func NewInternalServerError(message string, displayMessage string, err error) RestError {
	result := restErr{
		ErrMessage:     message,
		DisplayMessage: displayMessage,
		ErrStatus:      http.StatusInternalServerError,
		ErrError:       "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}

func NewTooManyRequestsError(message string, displayMessage string) RestError {
	return restErr{
		ErrMessage:     message,
		DisplayMessage: displayMessage,
		ErrStatus:      http.StatusTooManyRequests,
		ErrError:       "too_many_requests",
	}
}
