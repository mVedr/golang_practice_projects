package main

import (
	"context"
	"fmt"
	"log"

	userpb "github.com/mVedr/grpc_tut/gen/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()
	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Uuid: "hello",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", res)
}
