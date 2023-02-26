package customers

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomerByID(t *testing.T) {
	db, err := sqlx.Open("postgres", "user=postgres password=postgrespw dbname=car_rental sslmode=disable")
	assert.Nil(t, err)
	defer db.Close()
	t.Run("valid-uuid", func(t *testing.T) {
		var (
			id   = "a787151f-0116-4a79-8e1a-6f9c4b46af31"
			repo = New(db)
		)
		_, err := repo.GetByID(id)
		assert.Nil(t, err)
	})
	t.Run("invalid-uuid", func(t *testing.T) {
		var (
			id   = "invalid-uuid"
			repo = New(db)
		)
		_, err := repo.GetByID(id)
		assert.NotNil(t, err)
	})
}
