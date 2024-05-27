package server

import (
    "context"
    pb "user-service-module/proto/user/userpb"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
    s := NewUserServer()

    req := &pb.GetUserRequest{Id: 1}
    resp, err := s.GetUser(context.Background(), req)

    assert.NoError(t, err)
    assert.Equal(t, int32(1), resp.User.Id)
}

func TestListUsers(t *testing.T) {
    s := NewUserServer()

    req := &pb.ListUsersRequest{Ids: []int32{1, 2}}
    resp, err := s.ListUsers(context.Background(), req)

    assert.NoError(t, err)
    assert.Len(t, resp.Users, 2)
}

func TestSearchUsers(t *testing.T) {
    s := NewUserServer()

    req := &pb.SearchUsersRequest{City: "LA"}
    resp, err := s.SearchUsers(context.Background(), req)

    assert.NoError(t, err)
    assert.NotEmpty(t, resp.Users)
}
