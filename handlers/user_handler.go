package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"restfiber/constants"
	"restfiber/database"
	"restfiber/models"
	"time"
)

func Register(c *fiber.Ctx) error {
	var registeredUser models.User
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	username := data["username"]

	database.DB.Where("username = ?", username).First(&registeredUser)

	if registeredUser.Id != 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User is already registered",
		})
	}

	var newUser models.User
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	newUser.Username = username
	newUser.Password = string(password)

	if data["role"] == "ADMIN" {
		newUser.Role = "ADMIN"
	}

	database.DB.Create(&newUser)

	return c.JSON(newUser)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("username = ?", data["username"]).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	t, err := token.SignedString(constants.JwtSecret)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
