package product

import (
	"final_project/internal/infrastructure/container"
	productcontroller "final_project/internal/pkg/product/controller"
	productrepository "final_project/internal/pkg/product/repository"
	productusecase "final_project/internal/pkg/product/usecase"
	"final_project/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(r fiber.Router, conf *container.Container) {
	repository := productrepository.NewProductRepository(conf.Mysqldb)
	usecase := productusecase.NewProductUseCase(repository)
	controller := productcontroller.NewProductController(usecase)

	r.Post("", utils.AuthMiddleware, utils.HandleMultipleFile, controller.CreateProduct)
	r.Get("", controller.GetProducts)
	r.Get(":product_id", controller.GetProductByID)
	r.Put(":product_id", utils.AuthMiddleware, utils.HandleMultipleFile, controller.UpdateProductByID)
	r.Delete(":product_id", utils.AuthMiddleware, controller.DeleteProductByID)
}
