package server

import (
	"context"
	"fmt"
	"reflect"
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
				IsMarried: pb.MaritalStatus_MARRIED,
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
        missingIDs    []uint32
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
					IsMarried: pb.MaritalStatus_MARRIED,
				},
				{
					Id:      2,
					Fname:   "Bob",
					City:    "NY",
					Phone:   "9876543210",
					Height:  6.1,
					IsMarried: pb.MaritalStatus_SINGLE,
				},
			},
			expectedCode: 200,
			expectedErr:  nil,
			invalidIDs:   []uint32{},
		},
        {
            name:          "should return error for invalid ID",
            ids:           []uint32{0},
            expectedUsers: []*pb.User{},
            expectedCode:  400,
            expectedErr:   errors.ErrInvalidID,
            invalidIDs:    []uint32{0},
        },
		{
			name:          "should not list users with mix present and non-existent IDs",
			ids:           []uint32{1, 999},
			expectedUsers: []*pb.User{},
			expectedCode:  404,
			expectedErr:   errors.ErrUserNotFound,
			missingIDs:    []uint32{999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := userServer.ListUsers(context.Background(), &pb.ListUsersRequest{Ids: tt.ids})
			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			assert.Equal(t, tt.expectedUsers, resp.Users)
            fmt.Printf("Error: %v\n", err)
            fmt.Printf("Expected Error: %v\n", tt.expectedErr)
            fmt.Printf("Invalid IDs: %v\n", tt.invalidIDs)
			if tt.expectedErr != nil && len(tt.invalidIDs) != 0 {
				expectedError := fmt.Sprintf("%v: %v", errors.ErrInvalidID, tt.invalidIDs)
                if err.Error() != expectedError {
                    t.Errorf("ListUsers(%v) error = %v; want %v", tt.ids, err, expectedError)
                }
			} else if tt.expectedErr != nil && len(tt.invalidIDs) == 0{
                expectedError := fmt.Sprintf("%v: %v", errors.ErrUserNotFound, tt.missingIDs)
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
	userServer := NewUserServer()

	tests := []struct {
		name          string
		city          string
		phone         string
		married       pb.MaritalStatus
		expectedUsers []*pb.User
		expectedCode  uint32
		expectedErr   error
	}{
		{
			name: "should find users by city",
			city: "LA",
			expectedUsers: []*pb.User{
				{
					Id:     1,
					Fname:  "Steve",
					City:   "LA",
					Phone:  "9827329211",
					Height: 5.8,
					IsMarried: pb.MaritalStatus_MARRIED,
				},
				{
					Id:     3,
					Fname:  "Alice",
					City:   "LA",
					Phone:  "9876545876",
					Height: 5.5,
					IsMarried: pb.MaritalStatus_MARRIED,
				},
			},
			expectedCode: 200,
			expectedErr:  nil,
		},
		{
			name:  "should find users by phone",
			phone: "9876543210",
			expectedUsers: []*pb.User{
				{
					Id:     2,
					Fname:  "Bob",
					City:   "NY",
					Phone:  "9876543210",
					Height: 6.1,
					IsMarried: pb.MaritalStatus_SINGLE,
				},
			},
			expectedCode: 200,
			expectedErr:  nil,
		},
		{
			name:    "should find users by marital status",
			married: pb.MaritalStatus_MARRIED,
			expectedUsers: []*pb.User{
				{
					Id:     1,
					Fname:  "Steve",
					City:   "LA",
					Phone:  "9827329211",
					Height: 5.8,
					IsMarried: pb.MaritalStatus_MARRIED,
				},
				{
					Id:     3,
					Fname:  "Alice",
					City:   "LA",
					Phone:  "9876545876",
					Height: 5.5,
					IsMarried: pb.MaritalStatus_MARRIED,
				},
			},
			expectedCode: 200,
			expectedErr:  nil,
		},
		{
			name:          "should return error for invalid search request",
			city:          "InvalidCity",
			phone:         "0123456789",
			expectedUsers: []*pb.User{},
			expectedCode:  400,
			expectedErr:   errors.ErrInvalidFields,
		},
		{
			name:          "should return error when no users found",
			city:          "Unknown",
			expectedUsers: []*pb.User{},
			expectedCode:  404,
			expectedErr:   errors.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &pb.SearchUsersRequest{
				City:    tt.city,
				Phone:   tt.phone,
				IsMarried: tt.married,
			}
			resp, err := userServer.SearchUsers(context.Background(), req)
			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			reflect.DeepEqual(tt.expectedUsers, resp.Users)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
