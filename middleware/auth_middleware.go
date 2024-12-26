package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"restfiber/constants"
)

func AuthRequired(c *fiber.Ctx) error {
	tokeString := c.Get("Authorization")

	if tokeString == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokeString, func(token *jwt.Token) (interface{}, error) {
		return constants.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user_id", claims["user_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}

func GetUserIDFromContext(c *fiber.Ctx) uint {
	userId := c.Locals("user_id").(float64)

	return uint(userId)
}

func RoleRequired(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")

		for _, r := range requiredRoles {
			if r == role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access denied",
		})
	}
}
