package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
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
