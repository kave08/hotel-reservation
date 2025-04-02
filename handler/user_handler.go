package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kave08/hotel-reservation/db"
)

type UserHanlder struct {
	UserStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHanlder {
	return &UserHanlder{
		UserStore: userStore,
	}
}

func (h *UserHanlder) HandlePostUser(c fiber.Ctx) error {
	return nil
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
	users, err := h.UserStore.GetUsers(c.Context())
	if err != nil {
		return c.JSON(map[string]string{
			"error": fiber.ErrNotFound.Message,
		})
	}
	return c.JSON(users)
}
