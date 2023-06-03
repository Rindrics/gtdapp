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
	stuff := NewStuff(
		req.Title,
		req.Description,
	)

	repo := NewStuffRepository(s.DB)

	id, err := repo.Save(stuff)
	if err != nil {
		log.Fatalf("failed to save stuff: %v", err)
		return nil, err
	}

	stuff.Id = id

	return stuff, nil
}

func (s *InboxService) GetStuff(ctx context.Context, req *GetStuffRequest) (*Stuff, error) {
	repo := NewStuffRepository(s.DB)

	return repo.GetStuff(req.Id)
}

func (s *InboxService) GetStuffList(ctx context.Context, req *GetStuffListRequest) (*GetStuffListResponse, error) {
	repo := NewStuffRepository(s.DB)

	return repo.GetStuffList(req.Page, req.PerPage)
}

func (s *InboxService) UpdateStuff(ctx context.Context, req *UpdateStuffRequest) (*Stuff, error) {
	repo := NewStuffRepository(s.DB)

	return repo.UpdateStuff(req.Id, req.Title, req.Description)
}
