package main

import (
	"context"
	"log"
	"strconv"

	"github.com/imijanur/graphql-grpc-server/models"
	pb "github.com/imijanur/graphql-grpc-server/proto"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *server) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	// fetch users from the database get data from user_contacts and user_address tables as well
	users, err := models.Users(
		qm.Load("UserContacts"),
		qm.Load("UserAddresses"),
		qm.Limit(int(req.Limit)),
		qm.Offset(int(req.Offset)),
	).All(ctx, s.db)
	if err != nil {
		log.Printf("failed to fetch users: %v", err.Error())
		return nil, err
	}

	var pbUsers []*pb.CompleteUser
	for _, user := range users {
		var contact *pb.UserContact
		for _, cnt := range user.R.UserContacts {
			contact = &pb.UserContact{
				Id:        strconv.Itoa(int(cnt.ID)),
				FirstName: cnt.FirstName,
				LastName:  cnt.LastName,
				Phone:     cnt.Phone,
				UserId:    strconv.Itoa(int(cnt.UserID)),
			}
		}

		var addresses []*pb.UserAddress
		for _, address := range user.R.UserAddresses {
			addresses = append(addresses, &pb.UserAddress{
				Id:              strconv.Itoa(int(address.ID)),
				Name:            address.Name,
				Prefix:          address.Prefix.String,
				StreetAddress_1: address.StreetAddress1,
				StreetAddress_2: address.StreetAddress2.String,
				City:            address.City,
				State:           address.State,
				ZipCode:         address.ZipCode,
				UserId:          strconv.Itoa(int(address.UserID)),
			})
		}

		pbUsers = append(pbUsers, &pb.CompleteUser{
			User: &pb.User{
				Id:        strconv.Itoa(int(user.ID)),
				Email:     user.Email,
				Status:    user.Status,
				CreatedAt: user.CreatedAt.String(),
				UpdatedAt: user.ModifiedAt.String(),
			},
			Contact:   contact,
			Addresses: addresses,
		})
	}

	return &pb.GetUsersResponse{Users: pbUsers}, nil
}
