package domain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.banking/sardarmd/app/errs"
	"github.banking/sardarmd/logger"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type CustomerRepositoryDb struct {
	cosmoClient  *azcosmos.Client
	partitionKey *azcosmos.PartitionKey
	container    *azcosmos.ContainerClient
}

func (db *CustomerRepositoryDb) CreateDatabase() {

	database := azcosmos.DatabaseProperties{ID: "customerDb"}
	response, err := db.cosmoClient.CreateDatabase(context.Background(), database, nil)
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

	fmt.Printf("Container created. ActivityId %s", db.container.ID())
}

func (db *CustomerRepositoryDb) AddCustomerToDb() any {

	item := map[string]string{
		"id":          "88",
		"value":       `"name":"kasim", "value":"65" `,
		"customerKey": "/premium",
	}

	marshalled, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	itemResponse, err := db.container.CreateItem(context.Background(), *db.partitionKey, marshalled, nil)
	if err != nil {
		var responseErr *azcore.ResponseError
		errors.As(err, &responseErr)
		panic(responseErr)
	}

	fmt.Printf("Item created. ActivityId %s consuming %v RU", itemResponse.ActivityID, itemResponse.RequestCharge)
	handle(err)
	return nil

}

func (db CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppErrors) {

	itemResponse, err := db.container.ReadItem(context.Background(), *db.partitionKey, id, nil)
	if err != nil {

		var responseErr *azcore.ResponseError
		errors.As(err, &responseErr)

		if responseErr.ErrorCode == "NotFound" {
			logger.Error("Customer not found in Database")
			return nil, errs.NewNotFoundError()
		} else {
			logger.Error("Unexpected error occured in finding customer")
			return nil, errs.NewUnexpectedError()
		}

	}

	var customer Customer

	err = json.Unmarshal(itemResponse.Value, &customer)

	if err != nil {
		return nil, errs.NewUnexpectedError()

	}

	return &customer, nil

}

func (db CustomerRepositoryDb) FindAll() ([]Customer, error) {
	var customers []Customer

	queryPager := db.container.NewQueryItemsPager("select * from docs c", *db.partitionKey, nil)
	for queryPager.More() {
		queryResponse, err := queryPager.NextPage(context.Background())
		if err != nil {
			var responseErr *azcore.ResponseError
			errors.As(err, &responseErr)
			panic(responseErr)
		}

		for _, item := range queryResponse.Items {
			var customer Customer
			err = json.Unmarshal(item, &customer)
			customers = append(customers, customer)

			if err != nil {
				panic(err)
			}
		}

		fmt.Printf("Query page received with %v items. ActivityId %s consuming %v RU", len(queryResponse.Items), queryResponse.ActivityID, queryResponse.RequestCharge)

	}
	return customers, nil
}

func handle(err any) {

}
func NewCustomerRespositoryDb() CustomerRepositoryDb {
	var key = os.Getenv("AZURE_COSMOS_KEY")
	var endPoint = os.Getenv("AZURE_COSMOS_ENDPOINT")

	var Cred, err = azcosmos.NewKeyCredential(key)

	var Client, _ = azcosmos.NewClientWithKey(endPoint, Cred, nil)

	container, err := Client.NewContainer("customerDb", "customerContainer")

	if err != nil {
		panic(err)
	}

	pk := azcosmos.NewPartitionKeyString("/premium")

	return CustomerRepositoryDb{Client, &pk, container}

}
