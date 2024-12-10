package main

import (
	"log"

	"github.com/MuhammedAshifVnr/micro2/internal/delivery/http"
	"github.com/MuhammedAshifVnr/micro2/internal/queue"
	"github.com/MuhammedAshifVnr/micro2/internal/usecase"
	grpc "github.com/MuhammedAshifVnr/micro2/pkg/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Set up gRPC client
	userClient := grpc.NewUserClient("localhost:5001")

	// Initialize task queue for Method1
	taskQueue := queue.NewTaskQueue(10)
	defer taskQueue.Stop()

	// Initialize MethodService
	methodService := usecase.NewUserService(userClient, taskQueue)

	// Create Fiber app
	app := fiber.New()
	app.Use(logger.New())
	// Register routes
	handler := http.NewHandler(methodService,userClient.Client)
	http.RegisterRoutes(app, handler)

	// Start the server
	log.Println("Microservice2 running on port 8081")
	log.Fatal(app.Listen(":8081"))
}
