package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/imijanur/graphql-grpc-server/models"   // Update with the actual path to the generated models
	pb "github.com/imijanur/graphql-grpc-server/proto" // Update with the actual path to the generated protobuf package

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{}) // Update with your actual service server
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateAddress(t *testing.T) {
	// Create a mock database connection
	// Mock DB instance by sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}

	// Set up expectations for the mock database
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn) // Update with your actual service client

	// Create a fake user for testing
	user := &models.User{
		Email:      gofakeit.Email(),
		Status:     gofakeit.RandomString([]string{"active", "inactive"}),
		ModifiedAt: time.Now(),
	}

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	err = user.Insert(ctx, tx, boil.Infer())
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	req := &pb.CreateAddressRequest{
		UserId:         strconv.Itoa(user.ID),
		Name:           gofakeit.Company(),
		Prefix:         gofakeit.RandomString([]string{"Mr.", "Ms.", "Mrs.", "Dr."}),
		StreetAddress1: gofakeit.Street(),
		StreetAddress2: gofakeit.Street(),
		City:           gofakeit.City(),
		State:          gofakeit.State(),
		ZipCode:        gofakeit.Zip(),
	}

	resp, err := client.CreateAddress(ctx, req)
	if err != nil {
		t.Fatalf("CreateAddress failed: %v", err)
	}

	assert.NotNil(t, resp)
	assert.Equal(t, req.Name, resp.Address.Name)
	assert.Equal(t, req.Prefix, resp.Address.Prefix)
	assert.Equal(t, req.StreetAddress1, resp.Address.StreetAddress_1)
	assert.Equal(t, req.StreetAddress2, resp.Address.StreetAddress_2)
	assert.Equal(t, req.City, resp.Address.City)
	assert.Equal(t, req.State, resp.Address.State)
	assert.Equal(t, req.ZipCode, resp.Address.ZipCode)
	assert.Equal(t, req.UserId, resp.Address.UserId)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
