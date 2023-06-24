package usercontroller

import (
	"final_project/internal/helper"
	userdto "final_project/internal/pkg/user/dto"
	userusecase "final_project/internal/pkg/user/usecase"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetMyProfile(ctx *fiber.Ctx) error
	UpdateMyProfile(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	userusecase userusecase.UserUseCase
}

func NewUserController(userusecase userusecase.UserUseCase) UserController {
	return &UserControllerImpl{
		userusecase: userusecase,
	}
}

func (cn *UserControllerImpl) GetMyProfile(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals(("userID")))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	res, err := cn.userusecase.GetUserByID(c, userID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *UserControllerImpl) UpdateMyProfile(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	data := new(userdto.UserUpdate)
	if err := ctx.BodyParser(&data); err != nil {
		return helper.ResponseBuilder(*ctx, false, "Failed to parse request", err.Error(), nil, fiber.StatusBadRequest)
	}
	err := cn.userusecase.UpdateUserByID(c, userID, *data)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.PUTDATASUCCESS, nil, nil, fiber.StatusNoContent)
}
