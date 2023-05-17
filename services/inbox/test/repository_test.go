package test

import (
	"testing"

	"github.com/Rindrics/gtdapp-spec/services/inbox/internal"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	db := setupTestDB(t)

	repo := internal.NewStuffRepository(db)

	err := repo.Save(internal.NewStuff("Test title", "Test description"))

	assert.NoError(t, err, "failed to save stuff")
}
