package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "user-service-module/proto/user/userpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting gRPC client application...")

	// use grpc.Dial to connect to the running gRPC server
	// use grpc.WithBlock() to block until the connection is established
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
		return
	}

	defer conn.Close()

	// use the NewGreetingServiceClient method of the generated in the .pb.go file
	// to create a new GreetingServiceClient object
	// this object can be used to call methods implemented in the grpc server

	client := pb.NewUserServiceClient(conn)

	// Call GetUser
	getUser(client, 1)

	// Call ListUsers
	listUsers(client, []uint32{1, 2, 3})

	// Call SearchUsers
	searchUsers(client, "LA", "", pb.MaritalStatus_MARRIED)
}

func getUser(client pb.UserServiceClient, id uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetUserRequest{Id: id}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("GetUser Response: %v", res)
}

func listUsers(client pb.UserServiceClient, ids []uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ListUsersRequest{Ids: ids}
	res, err := client.ListUsers(ctx, req)
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	log.Printf("ListUsers Response: %v", res)
}

func searchUsers(client pb.UserServiceClient, city, phone string, isMarried pb.MaritalStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.SearchUsersRequest{City: city, Phone: phone, IsMarried: isMarried}
	res, err := client.SearchUsers(ctx, req)
	if err != nil {
		log.Fatalf("could not search users: %v", err)
	}
	log.Printf("SearchUsers Response: %v", res)
}
