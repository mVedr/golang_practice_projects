package main

import (
	"context"
	"log"
	"net"

	userpb "github.com/mVedr/grpc_tut/gen/go/user/v1"
	"google.golang.org/grpc"
)

type userService struct{}

func (u *userService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid:      req.Uuid,
			FullName:  "Ved",
			BirthYear: 2023,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("Failed to listen: %v \n", err)
	}
	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, &userService{})
	server.Serve(lis)
}
