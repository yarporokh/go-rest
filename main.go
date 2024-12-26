package main

import (
	"github.com/gofiber/fiber/v2"
	"restfiber/database"
	"restfiber/routes"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	routes.SetItemRoutes(app)
	routes.SetupUserRoutes(app)

	app.Listen(":3000")
}
