package app_errors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Status  int
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func InvalidRequestBodyError(errs error) *AppError {
	var errMsg string
	if validateErr, ok := errs.(validator.ValidationErrors); ok {
		errMsgs := make([]string, 0)
		for _, err := range validateErr {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"field: %s, value: %s, error: %s",
				err.Field(),
				err.Value(),
				err.Tag(),
			))
		}

		errMsg = strings.Join(errMsgs, " | ")
	} else {
		errMsg = errs.Error()
	}

	return &AppError{
		Status:  http.StatusBadRequest,
		Code:    "INVALID_REQUEST_BODY",
		Message: errMsg,
	}
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
