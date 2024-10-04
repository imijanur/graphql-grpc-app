package main

import (
	"context"
	"log"
	"strconv"

	"github.com/imijanur/graphql-grpc-server/models"
	pb "github.com/imijanur/graphql-grpc-server/proto"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *server) CreateAddress(ctx context.Context, req *pb.CreateAddressRequest) (*pb.CreateAddressResponse, error) {
	id, err := strconv.Atoi(req.UserId)
	if err != nil {
		log.Printf("failed to convert user id: %v", err.Error())
		return nil, err
	}

	address := &models.UserAddress{
		Name:           req.Name,
		Prefix:         null.StringFrom(req.Prefix),
		StreetAddress1: req.StreetAddress1,
		StreetAddress2: null.StringFrom(req.StreetAddress2),
		City:           req.City,
		State:          req.State,
		ZipCode:        req.ZipCode,
		UserID:         id,
	}

	err = address.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Printf("failed to insert address: %v", err.Error())
		return nil, err
	}

	return &pb.CreateAddressResponse{
		Address: &pb.UserAddress{
			Id:              strconv.Itoa(address.ID),
			Name:            address.Name,
			Prefix:          address.Prefix.String,
			StreetAddress_1: address.StreetAddress1,
			StreetAddress_2: address.StreetAddress2.String,
			City:            address.City,
			State:           address.State,
			ZipCode:         address.ZipCode,
			UserId:          req.UserId,
		},
	}, nil
}
