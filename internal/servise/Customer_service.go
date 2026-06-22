package servise

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/repository"
)

type CustomerServise struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerServise {
	return &CustomerServise{repo: repo}
}

func (s *CustomerServise) CreateCustomer(req model.RequestCustomer) (*model.Customer, error) {
	customer := &model.Customer{
		InnCustomer: req.InnCustomer,
		TypeCustomer: req.TypeCustomer,
		Firm: req.Firm,
		Name: req.Name,
		SecondName: req.SecondName,
		Surname: req.Surname,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := s.repo.Create(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerServise) GetAll() ([]model.Customer, error) {
	return s.repo.GetAll()
}

func (s *CustomerServise) GetAllByCustomer(customerID uint) ([]model.Customer, error) {
	return s.repo.GetAllByCustomer(customerID)
}