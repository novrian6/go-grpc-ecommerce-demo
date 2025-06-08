package main

import (
	"context"
	"log"
	"net"

	//pb "github.com/yourorg/proto-user/userpb"

	"google.golang.org/grpc"

	pb "example.com/go_grpc/ecommerce-demo/proto-user"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		Id:    req.Id,
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})
	log.Println("User Service running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
