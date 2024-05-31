package server

import (
	"context"
	"fmt"
	"testing"

	"user-service-module/internal/errors"
	pb "user-service-module/proto/user/userpb"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	userServer := NewUserServer()

	tests := []struct {
		name         string
		id           uint32
		expectedUser *pb.User
		expectedCode uint32
		expectedErr  error
	}{
		{
			name: "should get user by valid ID",
			id:   1,
			expectedUser: &pb.User{
				Id:      1,
				Fname:   "Steve",
				City:    "LA",
				Phone:   "9827329211",
				Height:  5.8,
				Married: true,
			},
			expectedCode: 200,
			expectedErr:  nil,
		},
		{
			name:         "should return error for invalid ID",
			id:           0,
			expectedUser: &pb.User{},
			expectedCode: 400,
			expectedErr:  errors.ErrInvalidID,
		},
		{
			name:         "should return error for non-existent user",
			id:           999,
			expectedUser: &pb.User{},
			expectedCode: 404,
			expectedErr:  errors.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := userServer.GetUser(context.Background(), &pb.GetUserRequest{Id: tt.id})
			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			assert.Equal(t, tt.expectedUser, resp.User)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestListUsers(t *testing.T) {
	userServer := NewUserServer()

	tests := []struct {
		name          string
		ids           []uint32
		expectedUsers []*pb.User
		expectedCode  uint32
		expectedErr   error
		invalidIDs    []uint32
	}{
		{
			name: "should list users by valid IDs",
			ids:  []uint32{1, 2},
			expectedUsers: []*pb.User{
				{
					Id:      1,
					Fname:   "Steve",
					City:    "LA",
					Phone:   "9827329211",
					Height:  5.8,
					Married: true,
				},
				{
					Id:      2,
					Fname:   "Bob",
					City:    "NY",
					Phone:   "9876543210",
					Height:  6.1,
					Married: false,
				},
			},
			expectedCode: 200,
			expectedErr:  nil,
			invalidIDs:   []uint32{},
		},
		{
			name:          "should not list users with mix present and non-existent IDs",
			ids:           []uint32{1, 999},
			expectedUsers: []*pb.User{},
			expectedCode:  404,
			expectedErr:   errors.ErrUserNotFound,
			invalidIDs:    []uint32{999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := userServer.ListUsers(context.Background(), &pb.ListUsersRequest{Ids: tt.ids})
			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			assert.Equal(t, tt.expectedUsers, resp.Users)
			if tt.expectedErr != nil && tt.invalidIDs != nil {
				expectedError := fmt.Sprintf("%v: %v", errors.ErrUserNotFound, tt.invalidIDs)
                if err.Error() != expectedError {
                    t.Errorf("ListUsers(%v) error = %v; want %v", tt.ids, err, expectedError)
                }
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSearchUsers(t *testing.T) {
    s := NewUserServer()

    req := &pb.SearchUsersRequest{City: "LA"}
    resp, err := s.SearchUsers(context.Background(), req)

    assert.NoError(t, err)
    assert.NotEmpty(t, resp.Users)
}
