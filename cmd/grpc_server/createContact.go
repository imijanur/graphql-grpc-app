package main

import (
	"context"
	"strconv"

	"github.com/imijanur/graphql-grpc-server/models"
	pb "github.com/imijanur/graphql-grpc-server/proto"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *server) CreateContact(ctx context.Context, req *pb.CreateContactRequest) (*pb.CreateContactResponse, error) {
	id, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}
	contact := &models.UserContact{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		UserID:    id,
	}

	err = contact.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &pb.CreateContactResponse{
		Contact: &pb.UserContact{
			Id:        strconv.Itoa(contact.ID),
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Phone:     contact.Phone,
			UserId:    req.UserId,
		},
	}, nil
}
