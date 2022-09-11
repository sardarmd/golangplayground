package service

import (
	"github.banking/sardarmd/app/domain"
	"github.banking/sardarmd/app/errs"
	"github.banking/sardarmd/dto"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppErrors)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {

	return s.repo.FindAll()

}
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppErrors) {

	c, error := s.repo.FindById(id)

	if error != nil {
		return nil, error
	} else {
		response := c.ToDto()
		return &response, nil
	}

}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
