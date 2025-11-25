package repository

import (
	"github.com/Baihaki123/customer-service/internal/domain"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(c *domain.Customer) error {
	return r.db.Create(c).Error
}

func (r *customerRepository) GetAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *customerRepository) GetByID(id int) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) Update(c *domain.Customer) error {
	return r.db.Save(c).Error
}

func (r *customerRepository) Delete(id int) error {
	return r.db.Delete(&domain.Customer{}, id).Error
}
