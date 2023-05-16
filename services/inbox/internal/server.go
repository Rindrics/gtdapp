package internal

import (
	"context"
	"database/sql"
	"log"
)

type InboxService struct {
	UnimplementedInboxServer
	DB *sql.DB
}

func (s *InboxService) Collect(ctx context.Context, req *CollectRequest) (*Stuff, error) {
	stuff := &Stuff{
		Item: &Item{
			Title:       req.Title,
			Description: req.Description,
		},
	}

	repo := NewStuffRepository(s.DB)

	err := repo.Save(stuff)
	if err != nil {
		log.Fatalf("failed to save stuff: %v", err)
		return nil, err
	}

	return stuff, nil
}
