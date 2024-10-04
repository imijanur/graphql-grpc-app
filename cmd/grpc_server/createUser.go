package main

import (
	"context"
	"strconv"

	"github.com/imijanur/graphql-grpc-server/models"
	pb "github.com/imijanur/graphql-grpc-server/proto"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &models.User{
		Email:  req.Email,
		Status: req.Status,
	}

	err := user.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:        strconv.Itoa(user.ID),
			Email:     user.Email,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.ModifiedAt.String(),
		},
	}, nil
}
