package main

import (
	"freedom/internal/user/handler"
	"log"
	"net"

	userpb "freedom/gen/proto/user"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	h := handler.NewUserHandler()
	userpb.RegisterUserServiceServer(s, h)

	log.Println("gRPC server listening at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
