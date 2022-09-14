package domain

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

type AccountRepository interface {
	CreateCustomer(account Account) (*Account, *errs.AppErrors)
}

func (c Account) ToAccountResponse() dto.NewAccountResponse {

	return dto.NewAccountResponse{
		Id: c.Id,
	}

}
