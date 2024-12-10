package client

import (
	"context"
	"log"

	pb "github.com/MuhammedAshifVnr/micro2/proto"
	"google.golang.org/grpc"
)

type UserClient struct {
	Client pb.UserServiceClient
}

func NewUserClient(address string) *UserClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	client := pb.NewUserServiceClient(conn)
	return &UserClient{Client: client}
}

func (uc *UserClient) ListUsers(ctx context.Context) ([]*pb.User, error) {
	resp, err := uc.Client.GetAllUsers(ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}
