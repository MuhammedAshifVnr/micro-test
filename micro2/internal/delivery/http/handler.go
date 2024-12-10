package http

import (
	"context"
	"fmt"

	"net/http"

	"github.com/MuhammedAshifVnr/micro2/internal/models"
	"github.com/MuhammedAshifVnr/micro2/internal/usecase"
	pb "github.com/MuhammedAshifVnr/micro2/proto"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userClient    pb.UserServiceClient
	methodService *usecase.UserService
}

func NewHandler(userService *usecase.UserService, user pb.UserServiceClient) *Handler {
	return &Handler{methodService: userService, userClient: user}
}

type MethodRequest struct {
	Method   int `json:"method"`
	WaitTime int `json:"waitTime"`
}

func (h *Handler) Methods(c *fiber.Ctx) error {
	var req MethodRequest
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("--=")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	ctx := context.Background()
	var userNames []string
	var err error

	switch req.Method {
	case 1:
		userNames, err = h.methodService.Method1(ctx, req.WaitTime)
	case 2:
		userNames, err = h.methodService.Method2(ctx, req.WaitTime)
	default:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid method"})
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"userNames": userNames})
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	res, err := h.userClient.CreateUser(context.Background(), &pb.CreateReq{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := h.userClient.GetUserByID(context.Background(), &pb.GetUserReq{Id: int64(id)})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.JSON(user)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	res, err := h.userClient.UpdateUser(context.Background(), &pb.UpdateReq{Id: uint64(id), Name: user.Name, Email: user.Email})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(res)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	res, err := h.userClient.DeleteUser(context.Background(), &pb.DeleteReq{Id: int64(id)})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusNoContent).JSON(res)
}
