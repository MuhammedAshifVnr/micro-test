package http

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, handler *Handler) {
	app.Post("/methods", handler.Methods)
	app.Post("/user", handler.CreateUser)
    app.Get("/user/:id", handler.GetUserByID)
    app.Put("/user/:id", handler.UpdateUser)
    app.Delete("/user/:id", handler.DeleteUser)
}
