package repo

import (
	"context"
	"encoding/json"

	"time"

	"github.com/MuhammedAshifVnr/micro1/internal/models"
	"github.com/go-redis/redis/v8"
)

type CacheRepository interface {
	SetUser(ctx context.Context, user *models.User, ttl time.Duration) error
	GetUser(ctx context.Context, id uint64) (*models.User, error)
	DeleteUser(ctx context.Context, id uint64) error
}

type CacheRepo struct {
	client *redis.Client
}

func NewCacheRepo(client *redis.Client) *CacheRepo {
	return &CacheRepo{client: client}
}

func (c *CacheRepo) SetUser(ctx context.Context, user *models.User, ttl time.Duration) error {
	data, _ := json.Marshal(user)
	return c.client.Set(ctx, c.userKey(uint64(user.ID)), data, ttl).Err()
}

func (c *CacheRepo) GetUser(ctx context.Context, id uint64) (*models.User, error) {
	data, err := c.client.Get(ctx, c.userKey(id)).Result()
	if err != nil {
		return nil, err
	}
	var user models.User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *CacheRepo) DeleteUser(ctx context.Context, id uint64) error {
	return c.client.Del(ctx, c.userKey(id)).Err()
}

func (c *CacheRepo) userKey(id uint64) string {
	return "user:" + string(rune(id))
}

