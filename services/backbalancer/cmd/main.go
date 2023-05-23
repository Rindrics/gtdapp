package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Rindrics/gtdapp-spec/services/backbalancer/internal"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	var opts []grpc.DialOption
	if strings.ToLower(os.Getenv("GRPC_INSECURE")) == "true" {
		opts = append(opts, grpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
		if err != nil {
			log.Fatalf("Failed to load TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	err := internal.RegisterInboxHandlerFromEndpoint(ctx, mux, "localhost:"+os.Getenv("GRPC_PORT_INBOX"), opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	port := os.Getenv("HTTP_PORT")
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
