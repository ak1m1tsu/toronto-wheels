package customers

import (
	"github.com/jmoiron/sqlx"
)

type Customer struct {
	ID            string `db:"id"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name"`
	Email         string `db:"email"`
	Phone         string `db:"phone"`
	Passhash      string `db:"passhash"`
	LicenseNumber string `db:"license_number"`
	Address       string `db:"address"`
}

type CustomerStorer interface {
	GetAll() ([]Customer, error)
	GetByID(string) (Customer, error)
	GetByEmail(string) (Customer, error)
	GetByPhone(string) (Customer, error)
	Insert(Customer) error
	Update(string, Customer) error
	Delete(string) error
}

type CustomerStore struct {
	*sqlx.DB
}

func New(db *sqlx.DB) *CustomerStore {
	return &CustomerStore{DB: db}
}

func (s *CustomerStore) GetAll() ([]Customer, error) {
	return []Customer{}, nil
}

func (s *CustomerStore) GetByID(id string) (Customer, error) {
	var (
		customer = Customer{}
		query    = "SELECT * FROM customers WHERE id=$1"
		err      = s.Get(&customer, query, id)
	)
	return customer, err
}

func (s *CustomerStore) GetByEmail(email string) (Customer, error) {
	return Customer{}, nil
}

func (s *CustomerStore) GetByPhone(phone string) (Customer, error) {
	return Customer{}, nil
}

func (s *CustomerStore) Insert(c Customer) error {
	return nil
}

func (s *CustomerStore) Update(id string, c Customer) error {
	return nil
}

func (s *CustomerStore) Delete(id string) error {
	return nil
}
