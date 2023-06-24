package store

import (
	"final_project/internal/infrastructure/container"
	storecontroller "final_project/internal/pkg/store/controller"
	storerepository "final_project/internal/pkg/store/repository"
	storeusecase "final_project/internal/pkg/store/usecase"
	"final_project/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func StoreRoute(r fiber.Router, conf *container.Container) {
	repository := storerepository.NewStoreRepository(conf.Mysqldb)
	usecase := storeusecase.NewStoreUseCase(repository)
	controller := storecontroller.NewStoreController(usecase)

	r.Get("", controller.GetAllStores)
	r.Get("/my", utils.AuthMiddleware, controller.GetMyStore)
	r.Put("/my", utils.HandleUploadFile, utils.AuthMiddleware, controller.UpdateStoreByID)
	r.Get("/:store_id", controller.GetStoreByID)
}
