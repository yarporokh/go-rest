package handlers

import (
	"github.com/gofiber/fiber/v2"
	"restfiber/database"
	"restfiber/middleware"
	"restfiber/models"
)

func GetItems(c *fiber.Ctx) error {
	var items []models.Item

	database.DB.Find(&items)

	return c.JSON(items)
}

func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")

	var item models.Item
	database.DB.First(&item, id)

	if item.Id == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(item)
}

func CreateItem(c *fiber.Ctx) error {
	var item models.Item

	if err := c.BodyParser(&item); err != nil {
		return err
	}

	userId := middleware.GetUserIDFromContext(c)
	item.UserId = userId

	database.DB.Create(&item)
	return c.JSON(item)
}

func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Item{}, id)
	return c.SendStatus(fiber.StatusOK)
}
