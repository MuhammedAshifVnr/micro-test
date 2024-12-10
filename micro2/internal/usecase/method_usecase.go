package usecase

import (
	"context"

	"time"

	"github.com/MuhammedAshifVnr/micro2/internal/queue"
	grpc "github.com/MuhammedAshifVnr/micro2/pkg/client"
)

type UserService struct {
	userClient *grpc.UserClient
	taskQueue  *queue.TaskQueue
}

func NewUserService(userClient *grpc.UserClient, taskQueue *queue.TaskQueue) *UserService {
	return &UserService{
		userClient: userClient,
		taskQueue:  taskQueue,
	}
}

// Method1: Sequential execution
func (s *UserService) Method1(ctx context.Context, waitTime int) ([]string, error) {
	result := make(chan []string)
	errChan := make(chan error)

	// Add the task to the sequential queue
	s.taskQueue.AddTask(func() {
		users, err := s.userClient.ListUsers(ctx)
		if err != nil {
			errChan <- err
			close(result)
			close(errChan)
			return
		}

		var userNames []string
		for _, user := range users {
			userNames = append(userNames, user.Name)
		}

		time.Sleep(time.Duration(waitTime) * time.Second) // Simulate work

		result <- userNames
		close(result)
		close(errChan)
	})

	// Retrieve results from the channels
	select {
	case res := <-result:
		return res, nil
	case err := <-errChan:
		return nil, err
	}
}

// Method2: Parallel execution
func (s *UserService) Method2(ctx context.Context, waitTime int) ([]string, error) {
	users, err := s.userClient.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var userNames []string
	for _, user := range users {
		userNames = append(userNames, user.Name)
	}

	time.Sleep(time.Duration(waitTime) * time.Second) // Simulate work
	return userNames, nil
}
