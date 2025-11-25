package usecase

import "github.com/Baihaki123/customer-service/internal/domain"

type CustomerUsecase interface {
	Create(c *domain.Customer) error
	GetAll() ([]domain.Customer, error)
	GetByID(id int) (*domain.Customer, error)
	Update(c *domain.Customer) error
	Delete(id int) error
}

type customerUsecase struct {
	repo domain.CustomerRepository
}

func NewCustomerUsecase(repo domain.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo: repo}
}

func (u *customerUsecase) Create(c *domain.Customer) error {
	// Validasi tambahan jika butuh bisa taruh di sini
	return u.repo.Create(c)
}

func (u *customerUsecase) GetAll() ([]domain.Customer, error) {
	return u.repo.GetAll()
}

func (u *customerUsecase) GetByID(id int) (*domain.Customer, error) {
	return u.repo.GetByID(id)
}

func (u *customerUsecase) Update(c *domain.Customer) error {
	// Biasanya ada validation, contoh:
	// if c.CstID == 0 { return errors.New("customer ID is required") }
	return u.repo.Update(c)
}

func (u *customerUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
