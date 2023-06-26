package transactioncontroller

import (
	"final_project/internal/helper"
	transactiondto "final_project/internal/pkg/transaction/dto"
	transactionusecase "final_project/internal/pkg/transaction/usecase"
	"log"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

type TransactionController interface {
	GetUserTransactions(ctx *fiber.Ctx) error
	GetUserTransactionByID(ctx *fiber.Ctx) error
	CreateTransaction(ctx *fiber.Ctx) error
}

type TransactionControllerImpl struct {
	transactionusecase transactionusecase.TransactionUseCase
}

func NewTransactionRepository(transactionusecase transactionusecase.TransactionUseCase) TransactionController {
	return &TransactionControllerImpl{
		transactionusecase: transactionusecase,
	}
}

func (cn *TransactionControllerImpl) GetUserTransactions(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, "UNAUTHORIZED", fiber.ErrUnauthorized.Message, nil, fiber.StatusUnauthorized)
	}
	res, err := cn.transactionusecase.GetUserTransactions(c, userID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *TransactionControllerImpl) GetUserTransactionByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, "UNAUTHORIZED", fiber.ErrUnauthorized.Message, nil, fiber.StatusUnauthorized)
	}
	transactionID := ctx.Params("transaction_id")
	if transactionID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	res, err := cn.transactionusecase.GetUserTransactionByID(c, userID, transactionID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *TransactionControllerImpl) CreateTransaction(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, "UNAUTHORIZED", fiber.ErrUnauthorized.Message, nil, fiber.StatusUnauthorized)
	}
	data := new(transactiondto.TransactionCreate)
	if err := ctx.BodyParser(&data); err != nil {
		log.Println(err)
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err, nil, fiber.StatusBadRequest)
	}
	data.UserID = userID
	res, err := cn.transactionusecase.CreateTransaction(c, *data)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.POSTDATASUCCESS, nil, res, fiber.StatusCreated)
}
