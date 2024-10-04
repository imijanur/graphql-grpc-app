package internal

import (
	"log"
	"sync"

	pb "github.com/imijanur/graphql-grpc-server/proto"
	"google.golang.org/grpc"
)

var (
	clientInstance pb.UserServiceClient
	clientOnce     sync.Once
)

func GetUserServiceClient() pb.UserServiceClient {
	clientOnce.Do(func() {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		clientInstance = pb.NewUserServiceClient(conn)
	})
	return clientInstance
}
