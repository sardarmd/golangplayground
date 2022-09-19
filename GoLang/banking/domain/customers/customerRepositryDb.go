package customers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.banking/sardarmd/errs"
	"github.banking/sardarmd/logger"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type CustomerRepositoryDb struct {
	cosmoClient  *azcosmos.Client
	partitionKey *azcosmos.PartitionKey
	container    *azcosmos.ContainerClient
}

func (db CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppErrors) {

	itemResponse, err := db.container.ReadItem(context.Background(), *db.partitionKey, id, nil)
	if err != nil {

		var responseErr *azcore.ResponseError
		errors.As(err, &responseErr)

		if responseErr.ErrorCode == "NotFound" {
			logger.Error("Customer not found in Database for id: " + id)
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

func NewCustomerRespositoryDb(client azcosmos.Client, partitionKey azcosmos.PartitionKey, container azcosmos.ContainerClient) CustomerRepositoryDb {

	return CustomerRepositoryDb{&client, &partitionKey, &container}

}
