package internal

import (
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
)

func NewStuff(title string, description string) *Stuff {
	timestamp, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		log.Fatalf("failed to create timestamp: %v", err)
	}

	return &Stuff{
		Item: &Item{
			Title:       title,
			Description: description,
			CreatedAt:   timestamp,
			UpdatedAt:   timestamp,
		},
	}
}
