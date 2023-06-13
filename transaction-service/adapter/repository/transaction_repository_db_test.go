package repository

import (
	"os"
	"testing"

	"github.com/ksd-mx/mx-payments/adapter/repository/fixture"
	"github.com/ksd-mx/mx-payments/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)
	repository := NewTransactionRepositoryDb(db)
	err := repository.SaveTransaction("1", "1", 100, entity.APPROVED, "")
	assert.Nil(t, err)
}
