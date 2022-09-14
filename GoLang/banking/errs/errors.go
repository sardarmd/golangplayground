package errs

import "net/http"

type AppErrors struct {
	Message string `json:"message"`
	Code    int    `json:",omitempty"`
}

func (e AppErrors) AsMessage() *AppErrors {
	return &AppErrors{Message: e.Message}
}
func NewNotFoundError() *AppErrors {
	return &AppErrors{"Data not found", http.StatusNotFound}
}

func NewUnexpectedError() *AppErrors {
	return &AppErrors{"Unexpected server error", http.StatusInternalServerError}
}

func NewCustomerException() *AppErrors {
	return &AppErrors{"Not able to create the account", http.StatusInternalServerError}
}

func UnsupportFormatException() *AppErrors {
	return &AppErrors{"Invalid input data, server is not able to process", http.StatusUnprocessableEntity}
}

func NewValidationError(msg string) *AppErrors {
	return &AppErrors{msg, http.StatusUnprocessableEntity}
}
