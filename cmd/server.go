package cmd

import (
    "user-service-module/internal/server"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "user-service-module/proto/user/userpb"
)

func main() {
    lis, err := net.Listen("tcp", ":33001")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, server.NewUserServer())

    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
