package repo

import (
	"context"

	"github.com/MuhammedAshifVnr/micro1/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uint64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint64) error
	ListAllUsers(ctx context.Context) (*[]models.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{
		db: DB,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) GetUserByID(ctx context.Context, ID uint64) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, ID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).
    Model(&models.User{}).
    Where("id = ?", user.ID).  
    Update("name", user.Name).  
    Error

}

func (r *UserRepo) DeleteUser(ctx context.Context, ID uint64) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, ID).Error
}

func (r *UserRepo) ListAllUsers(ctx context.Context) (*[]models.User, error) {
	var users []models.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
