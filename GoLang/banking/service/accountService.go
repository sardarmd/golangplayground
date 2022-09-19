package service

import (
	"time"

	"github.banking/sardarmd/domain/accounts"
	"github.banking/sardarmd/dto"
	"github.banking/sardarmd/errs"
)

type AccountService interface {
	AddCustomer(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppErrors)
}

type DefaultAccountService struct {
	repo accounts.AccountRepoistryDb
}

func (s DefaultAccountService) AddCustomer(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppErrors) {

	account := accounts.Account{
		Id:          time.Now().Local().String(),
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2016-01-02 14:02:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
		CustomerKey: "/premium",
	}
	c, error := s.repo.CreateCustomer(account)

	if error != nil {
		return nil, error
	} else {
		response := c.ToAccountResponse()
		return &response, nil
	}

}

func NewAccountService(repository accounts.AccountRepoistryDb) DefaultAccountService {
	return DefaultAccountService{repository}
}
