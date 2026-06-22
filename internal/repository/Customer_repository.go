package repository

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"gorm.io/gorm"
)

type CustomerRepository struct{
	db *gorm.DB
	customers []*model.Customer
	counter uint
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository{
	return &CustomerRepository{
		db: db,
		customers: make([]*model.Customer, 0),
		counter: 0,
	}
}

func (r *CustomerRepository) Create(customer *model.Customer) error{
	r.counter++
	customer.ID = r.counter
	r.customers = append(r.customers, customer)
	return nil
}

func (r *CustomerRepository) GetAll() ([]model.Customer, error) {
	result := make([]model.Customer, len(r.customers))
	for i, c := range r.customers {
		result[i] = *c
	}
	return result, nil
}

func (r *CustomerRepository) Delete(id uint) error {
	for i, c := range r.customers {
		if c.ID == id {
			r.customers = append(r.customers[:i], r.customers[i+1:]...)
			return nil
		}
	}
	return gorm.ErrRecordNotFound
}

func (r *CustomerRepository) GetAllByCustomer(CustomerID uint) ([]model.Customer, error){
	return []model.Customer{}, nil
}
