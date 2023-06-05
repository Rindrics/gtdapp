package test

import (
	"fmt"
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

func TestGetSavedStuff(t *testing.T) {
	db := setupTestDB(t)

	repo := internal.NewStuffRepository(db)

	id, err := repo.Save(internal.NewStuff("Test title", "Test description"))

	gotStuff, err := repo.GetStuff(id)
	if err != nil {
		t.Fatalf("failed to get stuff: %v", err)
	}

	assert.Equal(t, id, gotStuff.Id)
}

func TestGetSavedStuffList(t *testing.T) {
	db := setupTestDB(t)

	repo := internal.NewStuffRepository(db)

	// Create 13 stuffs
	for i := 0; i < 13; i++ {
		_, err := repo.Save(internal.NewStuff(fmt.Sprintf("Test title %d", i+1), "Test description"))
		if err != nil {
			t.Fatalf("failed to save stuff: %v", err)
		}
	}

	// get up to 5 each
	t.Run("can get first page", func(t *testing.T) {

		gotStuffList, err := repo.GetStuffList(1, 5)
		if err != nil {
			t.Fatalf("failed to get stuff list: %v", err)
		}

		assert.Equal(t, 5, len(gotStuffList.Stuffs))

		for i, stuff := range gotStuffList.Stuffs {
			assert.Equal(t, int64(i+1), stuff.Id)
		}
	})

	t.Run("can get second page", func(t *testing.T) {

		gotStuffList, err := repo.GetStuffList(2, 5)
		if err != nil {
			t.Fatalf("failed to get stuff list: %v", err)
		}

		assert.Equal(t, 5, len(gotStuffList.Stuffs))

		for i, stuff := range gotStuffList.Stuffs {
			assert.Equal(t, int64(i+6), stuff.Id)
		}
	})

	t.Run("can get the last page", func(t *testing.T) {

		gotStuffList, err := repo.GetStuffList(3, 5)
		if err != nil {
			t.Fatalf("failed to get stuff list: %v", err)
		}

		assert.Equal(t, 3, len(gotStuffList.Stuffs))

		for i, stuff := range gotStuffList.Stuffs {
			assert.Equal(t, int64(i+11), stuff.Id)
		}
	})
}

func TestUpdateSavedStuff(t *testing.T) {
	db := setupTestDB(t)
	repo := internal.NewStuffRepository(db)

	id, err := repo.Save(internal.NewStuff("Test title", "Test description"))
	if err != nil {
		t.Fatalf("failed to save stuff: %v", err)
	}

	_, err = repo.UpdateStuff(
		int64(id),
		"Updated title",
		"Updated description",
	)
	if err != nil {
		t.Fatalf("failed to update stuff: %v", err)
	}

	updatedStuff, err := repo.GetStuff(id)
	if err != nil {
		t.Fatalf("failed to get stuff: %v", err)
	}

	assert.Equal(t, "Updated title", updatedStuff.Item.Title)
	assert.Equal(t, "Updated description", updatedStuff.Item.Description)
}

func TestDeleteSavedStuff(t *testing.T) {
	db := setupTestDB(t)
	repo := internal.NewStuffRepository(db)

	id, err := repo.Save(internal.NewStuff("Test title", "Test description"))
	if err != nil {
		t.Fatalf("failed to save stuff: %v", err)
	}

	_, err = repo.DeleteStuff(
		int64(id),
	)
	if err != nil {
		t.Fatalf("failed to delete stuff: %v", err)
	}

	_, err = repo.GetStuff(id)
	if err == nil {
		t.Fatalf("stuff should be deleted")
	}

	assert.Equal(t, "stuff not found", err.Error())
}
