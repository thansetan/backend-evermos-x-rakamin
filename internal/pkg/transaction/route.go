package transaction

import (
	"final_project/internal/infrastructure/container"
	productrepository "final_project/internal/pkg/product/repository"
	transactioncontroller "final_project/internal/pkg/transaction/controller"
	transactionrepository "final_project/internal/pkg/transaction/repository"
	transactionusecase "final_project/internal/pkg/transaction/usecase"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(r fiber.Router, conf *container.Container) {
	trxRepo := transactionrepository.NewTransactionRepository(conf.Mysqldb)
	productRepo := productrepository.NewProductRepository(conf.Mysqldb)
	usecase := transactionusecase.NewTransactionUseCase(trxRepo, productRepo, conf.Mysqldb)
	controller := transactioncontroller.NewTransactionRepository(usecase)

	r.Get("", controller.GetUserTransactions)
	r.Get(":transaction_id", controller.GetUserTransactionByID)
	r.Post("", controller.CreateTransaction)
}
