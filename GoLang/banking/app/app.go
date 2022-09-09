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
	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRespositoryStub())}
	ch1 := CustomerHandler{service.NewCustomerService(domain.NewCustomerRespositoryDb())}
	router := mux.NewRouter()
	router.HandleFunc("/getcustomers", ch1.getAllCustomer).Methods(http.MethodGet)
	http.ListenAndServe("localhost:5057", router)
}
