package user

import (
	"final_project/internal/infrastructure/container"
	provincecityrepository "final_project/internal/pkg/provincecity/repository"
	usercontroller "final_project/internal/pkg/user/controller"
	userrepository "final_project/internal/pkg/user/repository"
	userusecase "final_project/internal/pkg/user/usecase"
	"final_project/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router, conf *container.Container) {
	userRepo := userrepository.NewUserRepository(conf.Mysqldb)
	provinceCityRepo := provincecityrepository.NewProvinceCityRepository()
	usecase := userusecase.NewUserUseCase(userRepo, provinceCityRepo)
	controller := usercontroller.NewUserController(usecase)

	r.Get("", utils.AuthMiddleware, controller.GetMyProfile)
	r.Put("", utils.AuthMiddleware, controller.UpdateMyProfile)
}
