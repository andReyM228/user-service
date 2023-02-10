package cars

import (
	"github.com/gofiber/fiber/v2"
	"user_service/repository/cars"
)

type Handler struct {
	carRepo cars.Repository
}

func NewHandler(repo cars.Repository) Handler {
	return Handler{
		carRepo: repo,
	}
}

func (h Handler) Get(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Create(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Delete(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Update(ctx *fiber.Ctx) error {
	return nil
}
