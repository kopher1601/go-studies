package domain

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

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
	}
}
