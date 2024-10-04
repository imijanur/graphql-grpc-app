package main

import (
	"database/sql"
	"log"
	"net"
	"os"

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

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
