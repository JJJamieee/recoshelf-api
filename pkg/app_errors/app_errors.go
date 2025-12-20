package app_errors

import (
	"net/http"
)

type AppError struct {
	Status  int
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// Common
var (
	InternalServerError = &AppError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Internal server error",
	}

	InvalidUserIDError = &AppError{
		Status:  http.StatusBadRequest,
		Code:    "INVALID_USER_ID",
		Message: "Invalid userID",
	}
)
