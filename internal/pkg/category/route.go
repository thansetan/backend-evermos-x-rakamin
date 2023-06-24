package category

import (
	"final_project/internal/infrastructure/container"
	categorycontroller "final_project/internal/pkg/category/controller"
	categoryrepository "final_project/internal/pkg/category/repository"
	categoryusecase "final_project/internal/pkg/category/usecase"
	"final_project/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoute(r fiber.Router, conf *container.Container) {
	repository := categoryrepository.NewCategoryRepository(conf.Mysqldb)
	usecase := categoryusecase.NewCategoryUseCase(repository)
	controller := categorycontroller.NewCategoryController(usecase)

	r.Get("", controller.GetCategories)               // don't think users have to be logged in to view categories
	r.Get(":category_id", controller.GetCategoryByID) // don't think users have to be logged in to view a category
	r.Post("", utils.AuthMiddleware, utils.CheckIsAdmin, controller.CreateCategory)
	r.Put(":category_id", utils.AuthMiddleware, utils.CheckIsAdmin, controller.UpdateCategoryByID)
	r.Delete(":category_id", utils.AuthMiddleware, utils.CheckIsAdmin, controller.DeleteCategoryByID)
}
