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
	return &AppErrors{"Data not found", http.StatusInternalServerError}
}
