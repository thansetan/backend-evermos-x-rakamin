package utils

import (
	"final_project/internal/helper"

	"github.com/gofiber/fiber/v2"
)

// @TODO : make middleware like Auth

func AuthMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("token")
	if token == "" {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	claims, err := DecodeJWT(token)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err, nil, fiber.StatusUnauthorized)
	}
	ctx.Locals("userID", claims["userID"])
	ctx.Locals("isAdmin", claims["isAdmin"])
	return ctx.Next()
}

func CheckIsAdmin(ctx *fiber.Ctx) error {
	isAdmin := ctx.Locals("isAdmin").(bool)
	if !isAdmin {
		return helper.ResponseBuilder(*ctx, false, "Failed to MODIFY data", fiber.ErrUnauthorized.Message, nil, fiber.StatusUnauthorized)
	}
	return ctx.Next()
}
