package internal

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type StuffRepository interface {
	Save(s *Stuff) (int64, error)
	GetStuff(id int64) (*Stuff, error)
	GetStuffList(page int64, per_page int64) (*GetStuffListResponse, error)
	UpdateStuff(id int64, title string, description string) (*Stuff, error)
	DeleteStuff(id int64) (*Stuff, error)
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
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("stuff not found")
		}
		log.Printf("failed to execute query: %v", err)
		return nil, err
	}

	s.Item.CreatedAt = timestamppb.New(createdAt)
	s.Item.UpdatedAt = timestamppb.New(updatedAt)

	return &s, nil
}

func (r *stuffRepository) GetStuffList(page int64, per_page int64) (*GetStuffListResponse, error) {
	stuffs := []*Stuff{}

	rows, err := r.db.Query("SELECT id, title, description, created_at, updated_at FROM stuff ORDER BY updated_at DESC LIMIT ? OFFSET ?", per_page, (page-1)*per_page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s Stuff
		s.Item = &Item{}
		var createdAt, updatedAt time.Time
		err = rows.Scan(&s.Id, &s.Item.Title, &s.Item.Description, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		s.Item.CreatedAt = timestamppb.New(createdAt)
		s.Item.UpdatedAt = timestamppb.New(updatedAt)
		stuffs = append(stuffs, &s)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &GetStuffListResponse{
		Stuffs: stuffs,
	}, nil
}

func (r *stuffRepository) UpdateStuff(id int64, title string, description string) (*Stuff, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatalf("failed to begin transaction: %v", err)
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE stuff SET title = ?, description = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v", err)
		return nil, err
	}

	now := time.Now()
	result, err := stmt.Exec(title, description, now.Format(time.RFC3339), id)
	if err != nil {
		log.Fatalf("failed to execute statement: %v", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows updated with id %d", id)
	}

	s := &Stuff{
		Id: id,
		Item: &Item{
			Title:       title,
			Description: description,
			UpdatedAt:   timestamppb.New(now),
		},
	}

	tx.Commit()

	return s, nil
}

func (r *stuffRepository) DeleteStuff(id int64) (*Stuff, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatalf("failed to begin transaction: %v", err)
		return nil, err
	}

	stmt, err := tx.Prepare("DELETE FROM stuff WHERE id = ?")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v", err)
		return nil, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		log.Fatalf("failed to execute statement: %v", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows deleted with id %d", id)
	}

	s := &Stuff{
		Id: id,
	}

	tx.Commit()

	return s, nil
}
