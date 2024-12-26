package routes

import (
	"github.com/gofiber/fiber/v2"
	"restfiber/handlers"
	"restfiber/middleware"
)

func SetItemRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/items", handlers.GetItems)
	app.Post("/items", middleware.AuthRequired, handlers.CreateItem)
	app.Get("/items/:id", handlers.GetItem)
	app.Delete("/items/:id", middleware.AuthRequired, middleware.RoleRequired("ADMIN"), handlers.DeleteItem)
}
