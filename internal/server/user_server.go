package server

import (
	"context"
	"errors"
	"net/http"
    "strings"

	pb "user-service-module/proto/user/userpb"
)

type UserServer struct {
    pb.UnimplementedUserServiceServer
    users []pb.User
}

func NewUserServer() *UserServer {
    // Sample data
    users := []pb.User{
        {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
        {Id: 2, Fname: "Bob", City: "NY", Phone: 9876543210, Height: 6.1, Married: false},
        {Id: 3, Fname: "Alice", City: "LA", Phone: 1928374650, Height: 5.5, Married: true},
    }

    return &UserServer{
        users: users,
    }
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

    for _, user := range s.users {
        if user.Id == req.Id {
            return &pb.GetUserResponse{
                StatusCode: http.StatusOK,
                User: &user,
            }, nil
        }
    }

    return &pb.GetUserResponse{
        StatusCode: http.StatusNotFound,
        User: &pb.User{},
    }, errors.New("user not found")
}

func (s *UserServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {

    var users []*pb.User

    for _, id := range req.Ids {
        for _, user := range s.users {
            if user.Id == id {
                users = append(users, &user)
            }
        }
    }
    return &pb.ListUsersResponse{
        StatusCode: http.StatusOK,
        Users: users,
    }, nil
}

func (s *UserServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
    
    var users []*pb.User
    
    for _, user := range s.users {
        if (req.City != "" && strings.EqualFold(user.City, req.City)) ||
            (req.Phone != 0 && user.Phone == req.Phone) ||
            (user.Married == req.Married) {
            users = append(users, &user)
        }
    }
    return &pb.SearchUsersResponse{
        StatusCode: http.StatusOK,
        Users: users,
    }, nil
}
