package test

import (
	"testing"

	"github.com/Rindrics/gtdapp-spec/services/inbox/internal"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	db := setupTestDB(t)

	repo := internal.NewStuffRepository(db)

	id, err := repo.Save(internal.NewStuff("Test title", "Test description"))

	assert.NoError(t, err, "failed to save stuff")
	assert.Equal(t, int64(1), id)
}

func TestGetStuff(t *testing.T) {
	db := setupTestDB(t)

	repo := internal.NewStuffRepository(db)

	id, err := repo.Save(internal.NewStuff("Test title", "Test description"))

	gotStuff, err := repo.GetStuff(id)
	if err != nil {
		t.Fatalf("failed to get stuff: %v", err)
	}

	assert.Equal(t, id, gotStuff.Id)
}
