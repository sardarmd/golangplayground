package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.banking/sardarmd/app/service"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	w.Header().Add("Content-type", "application/json")

	// CreateDatabase()
	// AddCustomerToDb()
	// ReadItem()
	
	GetItems()
	json.NewEncoder(w).Encode(customers)

}

func CreateDatabase() {

	key := os.Getenv("AZURE_COSMOS_KEY")
	endPoint := os.Getenv("AZURE_COSMOS_ENDPOINT")

	Cred, err := azcosmos.NewKeyCredential(key)

	Client, _ := azcosmos.NewClientWithKey(endPoint, Cred, nil)

	container, err := Client.NewContainer("customerDb", "customerContainer")

	database := azcosmos.DatabaseProperties{ID: "customerDb"}
	response, err := Client.CreateDatabase(context.Background(), database, nil)
	handle(err)

	fmt.Printf("Database created. ctivityId %s", response.ActivityID)

	if err != nil {
		var responseErr *azcore.ResponseError
		errors.As(err, &responseErr)
		panic(responseErr)
	}

	if err != nil {
		panic(err)
	}
	if err != nil {
		var responseErr *azcore.ResponseError
		errors.As(err, &responseErr)
		panic(responseErr)
	}

	fmt.Printf("Container created. ActivityId %s", container.ID())
}

func AddCustomerToDb() any {
	key := os.Getenv("AZURE_COSMOS_KEY")
	endPoint := os.Getenv("AZURE_COSMOS_ENDPOINT")

	Cred, err := azcosmos.NewKeyCredential(key)

	Client, _ := azcosmos.NewClientWithKey(endPoint, Cred, nil)

	container, err := Client.NewContainer("customerDb", "customerContainer")

	if err != nil {
		panic(err)
	}

	pk := azcosmos.NewPartitionKeyString("/premium")

	item := map[string]string{
		"id":          "88",
		"value":       `"name":"kasim", "value":"65" `,
		"customerKey": "/premium",
	}

	marshalled, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	itemResponse, err := container.CreateItem(context.Background(), pk, marshalled, nil)
	if err != nil {
		var responseErr *azcore.ResponseError
		errors.As(err, &responseErr)
		panic(responseErr)
	}

	fmt.Printf("Item created. ActivityId %s consuming %v RU", itemResponse.ActivityID, itemResponse.RequestCharge)
	handle(err)
	return nil

}

func ReadItem() {

	var key = os.Getenv("AZURE_COSMOS_KEY")
	var endPoint = os.Getenv("AZURE_COSMOS_ENDPOINT")

	var Cred, err = azcosmos.NewKeyCredential(key)

	var Client, _ = azcosmos.NewClientWithKey(endPoint, Cred, nil)

	container, err := Client.NewContainer("customerDb", "customerContainer")

	if err != nil {
		panic(err)
	}

	pk := azcosmos.NewPartitionKeyString("/premium")

	id := "2"
	itemResponse, err := container.ReadItem(context.Background(), pk, id, nil)
	if err != nil {
		var responseErr *azcore.ResponseError
		errors.As(err, &responseErr)
		panic(responseErr)
	}

	var itemResponseBody map[string]any

	err = json.Unmarshal(itemResponse.Value, &itemResponseBody)

	fmt.Printf(" Response value  %s", itemResponseBody["value"])

	if err != nil {
		panic(err)
	}

}

func GetItems() {

	var key = os.Getenv("AZURE_COSMOS_KEY")
	var endPoint = os.Getenv("AZURE_COSMOS_ENDPOINT")

	var Cred, err = azcosmos.NewKeyCredential(key)

	var Client, _ = azcosmos.NewClientWithKey(endPoint, Cred, nil)

	container, err := Client.NewContainer("customerDb", "customerContainer")
	if err != nil {
		panic(err)
	}

	pk := azcosmos.NewPartitionKeyString("/premium")

	queryPager := container.NewQueryItemsPager("select * from docs c", pk, nil)
	for queryPager.More() {
		queryResponse, err := queryPager.NextPage(context.Background())
		if err != nil {
			var responseErr *azcore.ResponseError
			errors.As(err, &responseErr)
			panic(responseErr)
		}

		for _, item := range queryResponse.Items {
			var itemResponseBody map[string]interface{}
			err = json.Unmarshal(item, &itemResponseBody)
			fmt.Printf(" Response value  %s", itemResponseBody["value"])
			if err != nil {
				panic(err)
			}
		}

		fmt.Printf("Query page received with %v items. ActivityId %s consuming %v RU", len(queryResponse.Items), queryResponse.ActivityID, queryResponse.RequestCharge)

	}
}

func handle(err any) {

}
