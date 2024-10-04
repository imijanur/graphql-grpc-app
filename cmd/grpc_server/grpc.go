package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	pb "github.com/imijanur/graphql-grpc-server/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

var (
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
)

func loadEnv() {
	// load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load the .env file: %v", err)
	}
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")
}

func main() {
	// load the .env file
	loadEnv()

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(recoveryInterceptor()),
	)
	pb.RegisterUserServiceServer(grpcServer, &server{db: db})

	// Channel to listen for OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// WaitGroup to wait for the server to shut down
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		<-sigChan
		log.Println("Received shutdown signal")

		// Create a context with a timeout to ensure the server shuts down within 5 seconds
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		grpcServer.GracefulStop()
		wg.Done()
	}()

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Wait for the server to shut down
	wg.Wait()
	log.Println("gRPC server has shut down gracefully")
}
