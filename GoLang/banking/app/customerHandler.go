package app

import (
	"encoding/json"
	"net/http"

	"github.banking/sardarmd/service"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, error := ch.service.GetCustomer(id)

	if error != nil {
		WriteResponse(w, error.Code, error.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customer)
	}

}

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
