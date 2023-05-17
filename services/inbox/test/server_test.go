package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Rindrics/gtdapp-spec/services/inbox/internal"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCollect(t *testing.T) {
	db, err := sql.Open("sqlite3", "./gtdapp.db")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	s := &internal.InboxService{DB: db}

	ctx := context.Background()
	req := &internal.CollectRequest{
		Title:       "Test title",
		Description: "Test description",
	}
	stuff, err := s.Collect(ctx, req)

	t.Run("Return is as expected", func(t *testing.T) {
		assert.NoError(t, err)

		assert.Equal(t, req.Title, stuff.Item.Title)
		assert.Equal(t, req.Description, stuff.Item.Description)
	})
}
