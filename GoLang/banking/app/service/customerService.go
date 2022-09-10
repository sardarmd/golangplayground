package service

import (
	"github.banking/sardarmd/app/domain"
	"github.banking/sardarmd/app/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppErrors)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {

	return s.repo.FindAll()

}
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppErrors) {

	return s.repo.FindById(id)

}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
