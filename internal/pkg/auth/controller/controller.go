package authcontroller

import (
	"final_project/internal/helper"
	authdto "final_project/internal/pkg/auth/dto"
	authusecase "final_project/internal/pkg/auth/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	authusecase authusecase.AuthUseCase
}

func NewAuthController(authusecase authusecase.AuthUseCase) AuthController {
	return &AuthControllerImpl{
		authusecase: authusecase,
	}
}

func (cn *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	c := ctx.Context()
	data := new(authdto.Register)
	if err := ctx.BodyParser(&data); err != nil {
		return helper.ResponseBuilder(*ctx, false, "Failed to parse request", err.Error(), nil, fiber.StatusBadRequest)
	}
	res, err := cn.authusecase.Register(c, *data)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, res, nil, *data, fiber.StatusCreated)
}

func (cn *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	c := ctx.Context()
	data := new(authdto.Login)
	if err := ctx.BodyParser(&data); err != nil {
		return helper.ResponseBuilder(*ctx, false, "Failed to parse request", err.Error(), nil, fiber.StatusBadRequest)
	}
	res, err := cn.authusecase.Login(c, *data)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.POSTDATASUCCESS, nil, res, fiber.StatusOK)
}
