package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/Rindrics/gtdapp-spec/services/inbox/internal"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCollect(t *testing.T) {
	db := setupTestDB(t)
	s := &internal.InboxService{DB: db}

	ctx := context.Background()
	req := &internal.CollectRequest{
		Title:       "Title from server_test.go",
		Description: "Description from server_test.go",
	}
	stuff, err := s.Collect(ctx, req)

	t.Run("Return is as expected", func(t *testing.T) {
		assert.NoError(t, err)

		assert.Equal(t, req.Title, stuff.Item.Title)
		assert.Equal(t, req.Description, stuff.Item.Description)
	})

	t.Run("Suff is saved", func(t *testing.T) {
		got, err := s.GetStuff(ctx, &internal.GetStuffRequest{
			Id: stuff.Id,
		})
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, int64(1), got.Id)
		log.Printf("server_test got: %v", got)
	})
}

func setupStuff(t *testing.T) *internal.InboxService {
	t.Helper()

	db := setupTestDB(t)
	s := &internal.InboxService{DB: db}

	ctx := context.Background()
	for i := 0; i < 13; i++ {
		req := &internal.CollectRequest{
			Title:       fmt.Sprintf("Title %d", i+1),
			Description: fmt.Sprintf("Description %d", i+1),
		}
		_, err := s.Collect(ctx, req)
		if err != nil {
			t.Fatalf("failed to save stuff: %v", err)
		}
	}

	return s
}

func TestGetCollectedStuff(t *testing.T) {
	s := setupStuff(t)

	ctx := context.Background()

	stuff, err := s.GetStuff(ctx, &internal.GetStuffRequest{
		Id: int64(2),
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, stuff.Item.Title, "Title 2")
	assert.Equal(t, stuff.Item.Description, "Description 2")
}

func TestGetCollectedStuffList(t *testing.T) {
	s := setupStuff(t)

	ctx := context.Background()

	stuffList, err := s.GetStuffList(ctx, &internal.GetStuffListRequest{
		Page:    int64(2),
		PerPage: int64(5),
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(stuffList.Stuffs), 5)

	expectedTitles := []string{"Title 6", "Title 7", "Title 8", "Title 9", "Title 10"}

	for i, title := range expectedTitles {
		assert.Equal(t, stuffList.Stuffs[i].Item.Title, title)
	}
}
