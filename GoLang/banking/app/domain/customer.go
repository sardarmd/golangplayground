package domain

import "github.banking/sardarmd/app/errs"

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zip_code    string `json:"zip_code"`
	DateofBirth string `json:"dob"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(string) (*Customer, *errs.AppErrors)
}
