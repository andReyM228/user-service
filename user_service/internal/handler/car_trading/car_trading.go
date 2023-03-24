package car_trading

import (
	"user_service/internal/service/car_trading"

	"github.com/gofiber/fiber/v2"

	"user_service/internal/handler"
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
	userID, err := ctx.ParamsInt("user_id")
	if err != nil {
		return handler.HandleError(ctx, err)
	}

	carID, err := ctx.ParamsInt("car_id")
	if err != nil {
		return handler.HandleError(ctx, err)
	}

	if err := h.carTrading.BuyCar(int64(userID), int64(carID)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
