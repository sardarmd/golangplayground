package domain

import (
	"github.banking/sardarmd/dto"
	"github.banking/sardarmd/errs"
)

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"full_name"`
	CustomerId  string `json:"customer_id"`
	City        string `json:"city"`
	Zip_code    string `json:"zip_code"`
	DateofBirth string `json:"dob"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(string) (*Customer, *errs.AppErrors)
}

func (c Customer) StatusAsText() string {
	if c.Status == "0" {
		return "inactive"
	} else {
		return "active"
	}

}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zip_code:    c.Zip_code,
		DateofBirth: c.DateofBirth,
		Status:      c.StatusAsText(),
	}

}
