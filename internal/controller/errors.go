package controller

import "net/http"

type SignalError interface {
	error
	Status() int
}

type HttpError struct {
	message string
	status  int
}

func NewHttpError(message string, status int) *HttpError {
	return &HttpError{
		message: message,
		status:  status,
	}
}

func (e *HttpError) Error() string {
	return e.message
}

func (e *HttpError) Status() int {
	return e.status
}

var (
	ErrInvalidID         = NewHttpError("invalid id", http.StatusBadRequest)
	ErrEmptyAuthHeader   = NewHttpError("authorization header is missing", http.StatusUnauthorized)
	ErrInvalidAuthHeader = NewHttpError("authorization header is invalid", http.StatusUnauthorized)
	ErrForbidden         = NewHttpError("you don't have access to this resource", http.StatusForbidden)
)
