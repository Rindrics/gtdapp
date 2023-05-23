package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/Rindrics/gtdapp-spec/services/inbox/internal"
	"github.com/Rindrics/gtdapp-spec/services/inbox/internal/health"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // We will run two servers concurrently

	// service endpoint on gRPC
	go func() {
		defer wg.Done()
		grpcServer := grpc.NewServer()

		dbDriver := os.Getenv("DB_DRIVER")
		dbSource := os.Getenv("DB_SOURCE")
		db, err := sql.Open(dbDriver, dbSource)
		if err != nil {
			log.Fatalf("gRPC server: failed to open database: %v", err)
		}

		inboxService := &internal.InboxService{
			DB: db,
		}
		internal.RegisterInboxServer(grpcServer, inboxService)
		reflection.Register(grpcServer)

		port := ":50051"
		listenPort, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("gRPC server: failed to listen: %v", err)
		}

		fmt.Printf("Server is starting on port %s\n", port)
		err = grpcServer.Serve(listenPort)
		if err != nil {
			log.Fatalf("gRPC server: failed to serve: %v", err)
		}
	}()

	// health endpoint on HTTP
	go func() {
		defer wg.Done()

		mux := http.NewServeMux()
		mux.HandleFunc("/health", health.HealthCheckHandler)

		port := ":8880"
		fmt.Printf("Serving health status on port %s\n", port)
		err := http.ListenAndServe(port, mux)
		if err != nil {
			log.Fatalf("HTTP server: failed to start HTTP server: %v", err)
		}
	}()

	wg.Wait()
}
