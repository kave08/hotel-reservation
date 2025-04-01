package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kave08/hotel-reservation/types"
)

func HandleGetUsers(c fiber.Ctx) error {
	u := types.User{
		ID:        "1",
		FirstName: "kave",
		LastName:  "hudj",
	}
	return c.JSON(u)
}

func HandleGetUser(c fiber.Ctx) error {
	return c.JSON(map[string]string{
		"user": "kaves",
	})
}
