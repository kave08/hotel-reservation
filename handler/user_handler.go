package handler

import (
	"net/http"

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

func (h *UserHanlder) HandlePostUser(c fiber.Ctx) error {
	var parms types.UserRequest

	if err := c.Bind().Body(&parms); err != nil {
		return c.JSON(map[string]any{
			"error":  fiber.ErrNotFound.Message,
			"status": http.StatusNotFound,
		})
	}

	user, err := types.NewUserParams(parms)
	if err != nil {
		return c.JSON(map[string]any{
			"error":  fiber.ErrBadRequest.Message,
			"status": http.StatusBadRequest,
		})
	}

	u, err := h.UserStore.CreateUser(c.Context(), user)
	if err != nil {
		return c.JSON(map[string]any{
			"error":  fiber.ErrInternalServerError.Message,
			"status": http.StatusInternalServerError,
		})
	}

	return c.JSON(u)
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
