package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s *CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			ID:          "1001",
			Name:        "Ashish",
			City:        "Tokyo",
			Zipcode:     "111-1111",
			DateOfBirth: "1999-09-09",
			Status:      "1",
		},
		{
			ID:          "1002",
			Name:        "Rob",
			City:        "Sapporo",
			Zipcode:     "111-1111",
			DateOfBirth: "1999-09-09",
			Status:      "1",
		},
	}
	return CustomerRepositoryStub{customers}
}
