package usecase

import (
	"context"
	"time"

	"github.com/MuhammedAshifVnr/micro1/internal/models"
	"github.com/MuhammedAshifVnr/micro1/internal/repo"
)

type UserService struct {
	userRepo  repo.UserRepository
	cacheRepo repo.CacheRepository
}

func NewUserService(userRepo repo.UserRepository, cacheRepo repo.CacheRepository) *UserService {
	return &UserService{userRepo: userRepo, cacheRepo: cacheRepo}
}

func (u *UserService) CreateUser(ctx context.Context, user *models.User) error {
	err := u.userRepo.CreateUser(ctx, user)
	if err == nil {
		u.cacheRepo.SetUser(ctx, user, time.Minute*10)
	}
	return err
}

func (u *UserService) GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	user, err := u.cacheRepo.GetUser(ctx, id)
	if err == nil {
		return user, nil
	}
	user, err = u.userRepo.GetUserByID(ctx, id)
	if err == nil {
		u.cacheRepo.SetUser(ctx, user, time.Minute*10)
	}
	return user, err
}

func (u *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	err := u.userRepo.UpdateUser(ctx, user)
	if err == nil {
		u.cacheRepo.SetUser(ctx, user, time.Minute*10)
	}
	return err
}

func (u *UserService) DeleteUser(ctx context.Context, id uint64) error {
	err := u.userRepo.DeleteUser(ctx, id)
	if err == nil {
		u.cacheRepo.DeleteUser(ctx, id)
	}
	return err
}

func (u *UserService)GetAllUsers(ctx context.Context)(*[]models.User,error){
	return u.userRepo.ListAllUsers(ctx)
}