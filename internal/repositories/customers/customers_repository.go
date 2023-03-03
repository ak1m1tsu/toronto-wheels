package customers

import (
	"github.com/jmoiron/sqlx"
	"github.com/romankravchuk/toronto-wheels/internal/repositories/models"
)

type CustomerStorer interface {
	FindAll() ([]models.Customer, error)
	FindById(string) (models.Customer, error)
	FindByEmail(string) (models.Customer, error)
	FindByPhone(string) (models.Customer, error)
	Insert(models.Customer) error
	Update(string, models.Customer) error
	Delete(string) error
}

type CustomerStore struct {
	*sqlx.DB
}

func New(db *sqlx.DB) *CustomerStore {
	return &CustomerStore{DB: db}
}

func (s *CustomerStore) FindAll(pg models.Pagination) ([]models.Customer, error) {
	var (
		customers []models.Customer
		query     = `
			SELECT * FROM customers
			ORDER BY $1
			OFFSET $2
			LIMIT $3
		`
		err = s.Select(&customers, query, pg.GetSort(), pg.GetOffSet(), pg.GetLimit())
	)
	return customers, err
}

func (s *CustomerStore) FindById(id string) (models.Customer, error) {
	var (
		customer = models.Customer{}
		query    = "SELECT * FROM customers WHERE id=$1"
		err      = s.Get(&customer, query, id)
	)
	return customer, err
}

func (s *CustomerStore) FindByEmail(email string) (models.Customer, error) {
	return models.Customer{}, nil
}

func (s *CustomerStore) FindByPhone(phone string) (models.Customer, error) {
	return models.Customer{}, nil
}

func (s *CustomerStore) Insert(c models.Customer) error {
	return nil
}

func (s *CustomerStore) Update(id string, c models.Customer) error {
	return nil
}

func (s *CustomerStore) Delete(id string) error {
	return nil
}
