package server

import (
	"context"
	"errors"
	"strings"
	pb "user-service-module/proto/user/userpb"
    "sync"
)

type UserServer struct {
    pb.UnimplementedUserServiceServer
    users []pb.User
    mu sync.Mutex
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
    s.mu.Lock()
    defer s.mu.Unlock()
    for _, user := range s.users {
        if user.Id == req.Id {
            return &pb.GetUserResponse{User: &user}, nil
        }
    }
    return nil, errors.New("user not found")
}

func (s *UserServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    var users []*pb.User

    for _, id := range req.Ids {
        for _, user := range s.users {
            if user.Id == id {
                users = append(users, &user)
            }
        }
    }
    return &pb.ListUsersResponse{Users: users}, nil
}

func (s *UserServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    var users []*pb.User
    
    for _, user := range s.users {
        if (req.City != "" && strings.EqualFold(user.City, req.City)) ||
            (req.Phone != 0 && user.Phone == req.Phone) ||
            (user.Married == req.Married) {
            users = append(users, &user)
        }
    }
    return &pb.SearchUsersResponse{Users: users}, nil
}
