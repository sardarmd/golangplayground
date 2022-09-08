package app

import (
	"fmt"
	"net/http"

	"github.banking/sardarmd/app/domain"
	"github.banking/sardarmd/app/service"
	"github.com/gorilla/mux"
)

func Start() {
	fmt.Printf("Starting new server")
	//wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRespositoryStub())}
	router := mux.NewRouter()
	router.HandleFunc("/getcustomers", ch.getAllCustomer).Methods(http.MethodGet)

	http.ListenAndServe("localhost:5056", router)
}
