package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Rindrics/gtdapp-spec/services/inbox/internal"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	db, err := sql.Open("sqlite3", "./gtdapp.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	inboxService := &internal.InboxService{
		DB: db,
	}
	internal.RegisterInboxServer(grpcServer, inboxService)

	listenPort, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = grpcServer.Serve(listenPort)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
