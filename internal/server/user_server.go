package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"user-service-module/internal/errors"
	"user-service-module/internal/utils"
	pb "user-service-module/proto/user/userpb"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	users map[uint32]pb.User
	mu    sync.Mutex
}

func NewUserServer() *UserServer {
	// Initialize the map with sample data
	// Using a map for faster lookups
	userMap := map[uint32]pb.User{
		1: {Id: 1, Fname: "Steve", City: "LA", Phone: "9827329211", Height: 5.8, Married: true},
		2: {Id: 2, Fname: "Bob", City: "NY", Phone: "9876543210", Height: 6.1, Married: false},
		3: {Id: 3, Fname: "Alice", City: "LA", Phone: "9876545876", Height: 5.5, Married: true},
	}

	return &UserServer{
		users: userMap,
	}
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !utils.IsIDValid(req.Id) {
		return &pb.GetUserResponse{
			StatusCode: http.StatusBadRequest,
			User:       &pb.User{},
		}, fmt.Errorf("%w: %d", errors.ErrInvalidID, req.Id)
	}
	if user, found := s.users[req.Id]; found {
		return &pb.GetUserResponse{
			StatusCode: http.StatusOK,
			User:       &user,
		}, nil
	}

	return &pb.GetUserResponse{
		StatusCode: http.StatusNotFound,
		User:       &pb.User{},
	}, fmt.Errorf("%w: %d", errors.ErrUserNotFound, req.Id)
}

func (s *UserServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var users []*pb.User = []*pb.User{}
	invalidIDs := utils.GetInvalidIDs(req.Ids)
	if len(invalidIDs) > 0 {
		return &pb.ListUsersResponse{
			StatusCode: http.StatusBadRequest,
			Users:      users,
		}, fmt.Errorf("%w: %v", errors.ErrInvalidID, invalidIDs)
	}

	usersNotFound := []uint32{}
	for _, id := range req.Ids {
		if user, found := s.users[id]; found {
			users = append(users, &user)
		} else {
			usersNotFound = append(usersNotFound, id)
		}
	}

	if len(usersNotFound) > 0 {
		return &pb.ListUsersResponse{
			StatusCode: http.StatusNotFound,
			Users:      []*pb.User{},
		}, fmt.Errorf("%w: %v", errors.ErrUserNotFound, usersNotFound)
	}

	return &pb.ListUsersResponse{
		StatusCode: http.StatusOK,
		Users:      users,
	}, nil
}

func (s *UserServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var users []*pb.User

    if isReqInvalid, err := utils.ValidateSearchRequest(req.City, req.Phone); !isReqInvalid {
        return &pb.SearchUsersResponse{
            StatusCode: http.StatusBadRequest,
            Users:      users,
        }, err
    }

	for _, user := range s.users {
		if (req.City != "" && strings.EqualFold(user.City, req.City)) ||
			(req.Phone != "" && user.Phone == req.Phone) ||
			(user.Married == req.Married) {
			users = append(users, &user)
		}
	}

    if len(users) == 0 {
        return &pb.SearchUsersResponse{
            StatusCode: http.StatusNotFound,
            Users:      users,
        }, fmt.Errorf("%w", errors.ErrUserNotFound)
    }

	return &pb.SearchUsersResponse{
		StatusCode: http.StatusOK,
		Users:      users,
	}, nil
}
