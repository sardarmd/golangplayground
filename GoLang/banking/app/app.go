package app

import (
	"net/http"
	"os"

	"github.banking/sardarmd/domain/accounts"
	"github.banking/sardarmd/domain/customers"

	"github.banking/sardarmd/service"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/gorilla/mux"
	"github.com/sardarmd/banking-lib/logger"
)

func Start() {

	logger.Info("Starting server")

	//wiring
	client, pk, container := getDbClient()

	ch := CustomerHandler{service.NewCustomerService(customers.NewCustomerRespositoryDb(*client, *pk, *container))}
	ah := AccountHandler{service.NewAccountService(accounts.NewAccountRespositoryDb(*client, *pk, *container))}

	router := mux.NewRouter()

	//Defining the Routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.createCustomer).Methods(http.MethodPost)

	//Listening to the port
	http.ListenAndServe("localhost:5059", router)
}

func getDbClient() (*azcosmos.Client, *azcosmos.PartitionKey, *azcosmos.ContainerClient) {
	var key = os.Getenv("AZURE_COSMOS_KEY")
	var endPoint = os.Getenv("AZURE_COSMOS_ENDPOINT")

	var Cred, err = azcosmos.NewKeyCredential(key)

	var Client, _ = azcosmos.NewClientWithKey(endPoint, Cred, nil)

	container, err := Client.NewContainer("customerDb", "customerContainer")

	if err != nil {
		panic(err)
	}

	pk := azcosmos.NewPartitionKeyString("/premium")

	return Client, &pk, container

}
