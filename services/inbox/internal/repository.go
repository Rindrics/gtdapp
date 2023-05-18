package internal

import (
	"database/sql"
	"log"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type StuffRepository interface {
	Save(s *Stuff) (int64, error)
	GetStuff(id int64) (*Stuff, error)
}

type stuffRepository struct {
	db *sql.DB
}

func NewStuffRepository(db *sql.DB) StuffRepository {
	return &stuffRepository{
		db: db,
	}
}

func (r *stuffRepository) Save(s *Stuff) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatalf("failed to begin transaction: %v", err)
		return -1, err
	}

	stmt, err := tx.Prepare("INSERT INTO stuff(title, description, created_at, updated_at) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v", err)
		return -1, err
	}

	result, err := stmt.Exec(s.Item.Title, s.Item.Description, s.Item.CreatedAt.AsTime().Format(time.RFC3339), s.Item.UpdatedAt.AsTime().Format(time.RFC3339))
	if err != nil {
		log.Fatalf("failed to execute statement: %v", err)
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("failed to get last insert id: %v", err)
	}

	tx.Commit()

	return id, nil
}

func (r *stuffRepository) GetStuff(id int64) (*Stuff, error) {
	s := Stuff{
		Item: &Item{},
	}

	row := r.db.QueryRow("SELECT id, title, description, created_at, updated_at FROM stuff WHERE id = ?", id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var createdAt, updatedAt time.Time
	err := row.Scan(&s.Id, &s.Item.Title, &s.Item.Description, &createdAt, &updatedAt)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return nil, err
	}

	s.Item.CreatedAt = timestamppb.New(createdAt)
	s.Item.UpdatedAt = timestamppb.New(updatedAt)

	return &s, nil
}
