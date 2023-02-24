package cars

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"user_service/internal/domain"
	"user_service/internal/handler"
	"user_service/internal/repository/cars"
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
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return handler.HandleError(ctx, err)
	}

	car, err := h.carRepo.Get(int64(id))
	if err != nil {
		return handler.HandleError(ctx, err)
	}

	payload, err := json.Marshal(car)
	if err != nil {
		return handler.HandleError(ctx, err)
	}

	return ctx.Send(payload)
}

func (h Handler) Update(ctx *fiber.Ctx) error {
	var car domain.Car
	if err := ctx.BodyParser(&car); err != nil {
		return handler.HandleError(ctx, err)
	}

	if err := h.carRepo.Update(car); err != nil {
		return handler.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Create(ctx *fiber.Ctx) error {
	var car domain.Car
	if err := ctx.BodyParser(&car); err != nil {
		return handler.HandleError(ctx, err)
	}

	if err := h.carRepo.Create(car); err != nil {
		return handler.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return handler.HandleError(ctx, err)
	}

	if err := h.carRepo.Delete(int64(id)); err != nil {
		return handler.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
