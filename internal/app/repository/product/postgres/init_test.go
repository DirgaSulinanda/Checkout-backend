package postgres

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mockDB, _, _ := sqlmock.New()
	db := sqlx.NewDb(mockDB, "postgres")

	repoMock := New(db)
	assert.NotNil(t, repoMock)
}
