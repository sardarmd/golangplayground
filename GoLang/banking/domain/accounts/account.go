package accounts

import (
	"github.banking/sardarmd/dto"
	"github.banking/sardarmd/errs"
)

type Account struct {
	Id          string `json:"id"`
	CustomerId  string `json:"customer_id"`
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
	CustomerKey string `json:"customerKey"`
}

//go:generate mockgen -destination=./mocks/domain/accounts/mockAccountRepository.go  -package=accounts github.banking/sardarmd/domain/accounts AccountRepository
type AccountRepository interface {
	CreateCustomer(account Account) (*Account, *errs.AppErrors)
}

func (c Account) ToAccountResponse() dto.NewAccountResponse {

	return dto.NewAccountResponse{
		Id: c.Id,
	}

}
