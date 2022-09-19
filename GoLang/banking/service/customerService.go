package service

import (
	"github.banking/sardarmd/domain/customers"
	domain "github.banking/sardarmd/domain/customers"
	"github.banking/sardarmd/dto"
	"github.banking/sardarmd/errs"
)

// go: genarate mockgen -destination=../mocks/service/mockCustomerService.go  -package= service github.banking/sardarmd/service CustomerService
type CustomerService interface {
	GetAllCustomer() ([]customers.Customer, error)
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
