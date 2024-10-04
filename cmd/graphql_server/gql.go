package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/imijanur/graphql-grpc-server/graph"
)

// Start HTTP Server (GraphQL or gRPC)
func startServer() {
	// Create a new GraphQL handler
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Create a new HTTP server with the GraphQL handler
	httpSrv := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	// Serve the GraphQL playground at the root path
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Starting GraphQL Server on :8080")

	// Start the server in a goroutine
	go func() {
		// log.Println("GraphQL Server goroutine on :8080")
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %s\n", err)
		}
		// log.Println("GraphQL Server started on :8080")
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	log.Println("Press Ctrl+C to shutdown server gracefully...")
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit

	log.Printf("Received signal: %s. Shutting down server...\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exited properly")
}

func main() {
	startServer()
}
