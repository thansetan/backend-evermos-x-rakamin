package auth

import (
	"final_project/internal/infrastructure/container"
	authcontroller "final_project/internal/pkg/auth/controller"
	authrepository "final_project/internal/pkg/auth/repository"
	authusecase "final_project/internal/pkg/auth/usecase"
	provincecityrepository "final_project/internal/pkg/provincecity/repository"
	storerepository "final_project/internal/pkg/store/repository"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, conf *container.Container) {
	authRepo := authrepository.NewAuthRepository(conf.Mysqldb)
	storeRepo := storerepository.NewStoreRepository(conf.Mysqldb)
	provinceCityRepo := provincecityrepository.NewProvinceCityRepository()
	usecase := authusecase.NewAuthUseCase(authRepo, storeRepo, provinceCityRepo, conf.Mysqldb)
	controller := authcontroller.NewAuthController(usecase)

	r.Post("/register", controller.Register)
	r.Post("/login", controller.Login)
}
