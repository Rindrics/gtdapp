package test

import (
	"testing"

	"github.com/Rindrics/gtdapp-spec/services/inbox/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewStuff(t *testing.T) {
	title := "Test title"
	description := "Test description"
	stuff := internal.NewStuff(title, description)

	assert.Equal(t, title, stuff.Item.Title)
	assert.Equal(t, description, stuff.Item.Description)
	assert.NotNil(t, stuff.Item.CreatedAt)
	assert.NotNil(t, stuff.Item.UpdatedAt)
}
