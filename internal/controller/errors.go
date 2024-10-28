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

func NewCustomError(message string, status int) *HttpError {
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
	ErrInvalidID         = NewCustomError("invalid id", http.StatusBadRequest)
	ErrEmptyAuthHeader   = NewCustomError("authorization header is missing", http.StatusUnauthorized)
	ErrInvalidAuthHeader = NewCustomError("authorization header is invalid", http.StatusUnauthorized)
	ErrForbidden         = NewCustomError("you don't have access to this resource", http.StatusForbidden)
)
