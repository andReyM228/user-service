package handler

import (
	"github.com/gofiber/fiber/v2"
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
	return nil
}

func (h Handler) Update(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Create(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Delete(ctx *fiber.Ctx) error {
	return nil
}
