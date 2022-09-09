package domain

type CustomerRepositoryStub struct {
	customers []Customer
	test      string
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRespositoryStub() CustomerRepositoryStub {

	customers := []Customer{{"1001", "Sardar", "Bangalore", "560067", "25-08-1987", "1"}, {"1002", "Sadia", "Kanpur", "201301", "25-08-1987", "1"}}

	return CustomerRepositoryStub{customers, ""}
}
