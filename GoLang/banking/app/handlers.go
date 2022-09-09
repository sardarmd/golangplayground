package app

import (
	"encoding/json"
	"net/http"

	"github.banking/sardarmd/app/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
