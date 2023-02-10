package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"user_service/domain"
	"user_service/repository"
)

type Handler struct {
	userRepo repository.Repository
}

func NewHandler(repo repository.Repository) Handler {
	return Handler{
		userRepo: repo,
	}
}

func (h Handler) Get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	user, err := h.userRepo.Get(int64(id))
	if err != nil {
		return err
	}

	payload, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return ctx.Send(payload)
}

func (h Handler) Update(ctx *fiber.Ctx) error {
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	if err := h.userRepo.Update(user); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Create(ctx *fiber.Ctx) error {
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	if err := h.userRepo.Create(user); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	if err := h.userRepo.Delete(int64(id)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
