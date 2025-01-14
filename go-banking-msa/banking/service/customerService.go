package service

import "banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepositoryStub
}

func (d *DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return d.repo.FindAll()
}

func NewCustomerService(repo domain.CustomerRepositoryStub) CustomerService {
	return &DefaultCustomerService{repo: repo}
}
