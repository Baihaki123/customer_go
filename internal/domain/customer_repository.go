package domain

type CustomerRepository interface {
	Create(c *Customer) error
	GetAll() ([]Customer, error)
	GetByID(id int) (*Customer, error)
	Update(c *Customer) error
	Delete(id int) error
}
