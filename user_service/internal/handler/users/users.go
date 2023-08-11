package users

import (
	"encoding/json"
	"github.com/andReyM228/lib/rabbit"

	"user_service/internal/domain"
	"user_service/internal/domain/errs"
	"user_service/internal/repository/users"
	users_service "user_service/internal/service/users"

	"github.com/andReyM228/lib/responder"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userRepo    users.Repository
	userService users_service.Service
	rabbit      rabbit.Rabbit
}

func NewHandler(repo users.Repository, service users_service.Service, rabbit rabbit.Rabbit) Handler {
	return Handler{
		userRepo:    repo,
		userService: service,
		rabbit:      rabbit,
	}
}

func (h Handler) Get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	user, err := h.userRepo.Get(domain.FieldID, id)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	payload, err := json.Marshal(user)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.Send(payload)
}

func (h Handler) Update(ctx *fiber.Ctx) error {
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.userRepo.Update(user); err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Create(ctx *fiber.Ctx) error {
	var user domain.User
	if err := ctx.BodyParser(&user); err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.userService.Registration(user); err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h Handler) BrokerCreate(request []byte) error {
	var req rabbit.RequestModel
	if err := json.Unmarshal(request, &req); err != nil {
		return err
	}

	var user domain.User
	if err := json.Unmarshal(req.Payload, &user); err != nil {
		return err
	}

	if err := h.userService.Registration(user); err != nil {
		return h.rabbit.Reply(req.ReplyTopic, 500, nil)
	}

	return h.rabbit.Reply(req.ReplyTopic, 200, nil)
}

func (h Handler) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	if err := h.userRepo.Delete(int64(id)); err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Login(ctx *fiber.Ctx) error {
	var request loginRequest
	if err := ctx.BodyParser(&request); err != nil {
		return responder.HandleError(ctx, err)
	}

	if request.ChatID == 0 || request.Password == "" {
		return responder.HandleError(ctx, errs.BadRequestError{Cause: "wrong body"})
	}

	userID, err := h.userService.Login(request.ChatID, request.Password)
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	payload, err := json.Marshal(loginResponse{userID})
	if err != nil {
		return responder.HandleError(ctx, err)
	}

	return ctx.Send(payload)
}
