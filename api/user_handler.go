package api

import "github.com/gofiber/fiber/v3"

func HandleGetUser(c fiber.Ctx) error {
	return c.JSON(map[string]string{
		"user": "kaves",
	})
}
