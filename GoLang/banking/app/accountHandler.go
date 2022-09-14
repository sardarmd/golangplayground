package app

import (
	"encoding/json"
	"net/http"

	"github.banking/sardarmd/dto"
	"github.banking/sardarmd/service"
	validator "github.banking/sardarmd/utils"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ch *AccountHandler) createCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {

		err := validator.ValidateNewAccount(request)

		if err != nil {
			WriteResponse(w, err.Code, err.AsMessage())
		} else {
			request.CustomerId = id
			customer, e := ch.service.AddCustomer(request)

			if e != nil {
				WriteResponse(w, e.Code, err.Message)
			} else {
				WriteResponse(w, http.StatusCreated, customer)
			}
		}

	}

}
