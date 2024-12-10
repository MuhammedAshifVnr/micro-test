package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/MuhammedAshifVnr/micro1/internal/models"
	"github.com/MuhammedAshifVnr/micro1/internal/usecase"
	pb "github.com/MuhammedAshifVnr/micro1/proto"
	"gorm.io/gorm"
)

// GRPCHandler represents the gRPC handler for user-related operations.
type GRPCHandler struct {
	userService                       usecase.UserService
	pb.UnimplementedUserServiceServer // Embed the gRPC service interface
}

// NewGRPCHandler creates a new GRPCHandler instance.
func NewGRPCHandler(userService usecase.UserService) *GRPCHandler {
	return &GRPCHandler{userService: userService}
}

// CreateUser handles the gRPC request to create a new user.
func (h *GRPCHandler) CreateUser(ctx context.Context, req *pb.CreateReq) (*pb.CommonRes, error) {
	err := h.userService.CreateUser(ctx, &models.User{Name: req.Name, Email: req.Email})
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return &pb.CommonRes{
		Message: "user created",
	}, nil
}

// GetUser handles the gRPC request to retrieve a user by ID.
func (h *GRPCHandler) GetUserByID(ctx context.Context, req *pb.GetUserReq) (*pb.User, error) {
	user, err := h.userService.GetUserByID(ctx, uint64(req.Id))
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}

	return &pb.User{
		Id:    uint64(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// ListUsers handles the gRPC request to list all users.
func (h *GRPCHandler) GetAllUsers(ctx context.Context, req *pb.EmptyRequest) (*pb.UserListResponse, error) {
	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		log.Printf("Error listing users: %v", err)
		return nil, err
	}

	var userList []*pb.User
	for _, user := range *users {
		userList = append(userList, &pb.User{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &pb.UserListResponse{
		Users: userList,
	}, nil
}

// UpdateUser handles the gRPC request to update user details.
func (h *GRPCHandler) UpdateUser(ctx context.Context, req *pb.UpdateReq) (*pb.CommonRes, error) {
	fmt.Println("fre", req.Id)
	err := h.userService.UpdateUser(ctx, &models.User{Name: req.Name, Email: req.Email, Model: gorm.Model{ID: uint(req.Id)}})
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	return &pb.CommonRes{
		Message: "user updated",
	}, nil
}

// DeleteUser handles the gRPC request to delete a user by ID.
func (h *GRPCHandler) DeleteUser(ctx context.Context, req *pb.DeleteReq) (*pb.CommonRes, error) {
	err := h.userService.DeleteUser(ctx, uint64(req.Id))
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return nil, err
	}

	return &pb.CommonRes{
		Message: "user deleted",
	}, nil
}
