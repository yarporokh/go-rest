package routes

import (
	"github.com/gofiber/fiber/v2"
	"restfiber/handlers"
)

func SetupUserRoutes(app *fiber.App) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
}
