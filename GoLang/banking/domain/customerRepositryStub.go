package domain

type CustomerRepositoryStub struct {
	customers []Customer
	customer  Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
func (s CustomerRepositoryStub) FindById(id string) (*Customer, error) {
	return &s.customer, nil
}

func NewCustomerRespositoryStub() CustomerRepositoryStub {

	customers := []Customer{{"1001", "Sardar", "", "Bangalore", "560067", "25-08-1987", "1"}, {"1002", "", "Sadia", "Kanpur", "201301", "25-08-1987", "1"}}

	custo := Customer{"1001", "Sardar", "", "Bangalore", "560067", "25-08-1987", "1"}
	return CustomerRepositoryStub{customers, custo}
}
