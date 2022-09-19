package accounts

import (
	"context"
	"encoding/json"

	"github.com/sardarmd/banking-lib/errs"
	"github.com/sardarmd/banking-lib/logger"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type AccountRepoistryDb struct {
	cosmoClient  *azcosmos.Client
	partitionKey *azcosmos.PartitionKey
	container    *azcosmos.ContainerClient
}

func (db AccountRepoistryDb) CreateCustomer(account Account) (*Account, *errs.AppErrors) {

	marshalled, err := json.Marshal(account)
	if err != nil {
		return nil, errs.UnsupportFormatException()
	}

	var acc Account
	itemResponse, err := db.container.CreateItem(context.Background(), *db.partitionKey, marshalled, nil)

	if err != nil {
		logger.Error("Exception while creating new account")
		return nil, errs.NewCustomerException()
	}

	logger.Info("New account created successfully  Activity Id" + itemResponse.ActivityID)

	oldId := account.Id
	account.Id = itemResponse.ActivityID
	acc.Id = itemResponse.ActivityID
	updatedAccount, err := json.Marshal(account)

	if err != nil {
		return nil, errs.UnsupportFormatException()
	}

	updateResponse, err := db.container.ReplaceItem(context.Background(), *db.partitionKey, oldId, updatedAccount, nil)

	if err != nil {
		logger.Error("Exception while updating after creating the account")
		return nil, errs.NewCustomerException()
	}
	logger.Info("Updated account successfully  Activity Id" + updateResponse.ActivityID)

	return &acc, nil

}

func NewAccountRespositoryDb(client azcosmos.Client, partitionKey azcosmos.PartitionKey, container azcosmos.ContainerClient) AccountRepoistryDb {

	return AccountRepoistryDb{&client, &partitionKey, &container}

}
