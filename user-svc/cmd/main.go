package main

import (
	"github.com/MuhammedAshifVnr/micro1/internal/config"
	"github.com/MuhammedAshifVnr/micro1/internal/repo"
	"github.com/MuhammedAshifVnr/micro1/internal/usecase"
)

func main() {
	config.LoadEnv()
	db, rdb := config.InitDB()

	userRepo := repo.NewUserRepo(db)
	redisClient := repo.NewCacheRepo(rdb)
	userService := usecase.NewUserService(userRepo, redisClient)
	config.RunGRPCServer(*userService)
}
