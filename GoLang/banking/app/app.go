package app

import (
	"net/http"

	"github.banking/sardarmd/app/domain"
	"github.banking/sardarmd/app/service"
	"github.banking/sardarmd/logger"
	"github.com/gorilla/mux"
)

func Start() {

	logger.Info("Starting server")
	//wiring
	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRespositoryStub())}
	ch1 := CustomerHandler{service.NewCustomerService(domain.NewCustomerRespositoryDb())}
	router := mux.NewRouter()

	//Defining the Routes
	router.HandleFunc("/getcustomers", ch1.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/getcustomers/{customer_id:[0-9]+}", ch1.getCustomer).Methods(http.MethodGet)

	//Listening to the port
	http.ListenAndServe("localhost:5058", router)
}
