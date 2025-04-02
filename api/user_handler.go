package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kave08/hotel-reservation/db"
	"github.com/kave08/hotel-reservation/types"
)

type UserHanlder struct {
	UserStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHanlder {
	return &UserHanlder{
		UserStore: userStore,
	}
}

func (h *UserHanlder) HandleGetUser(c fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.UserStore.GetUserByID(c.Context(), id)
	if err != nil {
		return c.JSON(map[string]string{
			"error": fiber.ErrNotFound.Message,
		})
	}

	return c.JSON(map[string]string{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	})
}

func (h *UserHanlder) HandleGetUsers(c fiber.Ctx) error {
	u := types.User{
		ID:        "1",
		FirstName: "kave",
		LastName:  "hudj",
	}
	return c.JSON(u)
}
