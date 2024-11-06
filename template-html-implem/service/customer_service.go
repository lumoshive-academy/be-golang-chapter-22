package service

import (
	"be-golang-chapter-22/template-html-implem/model"
	"be-golang-chapter-22/template-html-implem/repository"
)

type CustomerService struct {
	RepoCustomer repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return CustomerService{RepoCustomer: repo}
}

func (cs *CustomerService) LoginService(customer model.Customer) error {
	err := cs.RepoCustomer.Login(&customer)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CustomerService) CustomerByID(id int) (*model.Customer, error) {

	customer, err := cs.RepoCustomer.CustomerByID(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (cs *CustomerService) AllCustomer() (*[]model.Customer, error) {

	customer, err := cs.RepoCustomer.AllCustomer()
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (cs *CustomerService) Register(customer *model.Customer) error {
	err := cs.RepoCustomer.Register(customer)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CustomerService) CheckToken(token string) string {
	result := cs.RepoCustomer.CheckToken(token)
	// if result == "" {
	// 	return ""
	// }
	return result
}
