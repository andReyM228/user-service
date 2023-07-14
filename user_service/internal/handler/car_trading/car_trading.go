package car_trading

import (
	"user_service/internal/service/car_trading"

	"github.com/andReyM228/lib/responder"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	carTrading car_trading.Service
}

func NewHandler(carTrading car_trading.Service) Handler {
	return Handler{
		carTrading: carTrading,
	}
}

func (h Handler) BuyCar(ctx *fiber.Ctx) error {
	chatID, err := ctx.ParamsInt("chat_id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	carID, err := ctx.ParamsInt("car_id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.carTrading.BuyCar(int64(chatID), int64(carID)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) SellCar(ctx *fiber.Ctx) error {
	chatID, err := ctx.ParamsInt("chat_id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	carID, err := ctx.ParamsInt("car_id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.carTrading.SellCar(int64(chatID), int64(carID)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
