package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Rindrics/gtdapp-spec/services/inbox/internal"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	inboxService := &internal.InboxService{
		DB: db,
	}
	internal.RegisterInboxServer(grpcServer, inboxService)
	reflection.Register(grpcServer)

	port := ":50051"
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Server is starting on port %s\n", port)
	err = grpcServer.Serve(listenPort)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
