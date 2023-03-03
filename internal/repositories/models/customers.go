package models

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
