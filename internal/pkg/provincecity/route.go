package provincecity

import (
	provincecitycontroller "final_project/internal/pkg/provincecity/controller"
	provincecityrepository "final_project/internal/pkg/provincecity/repository"
	provincecityusecase "final_project/internal/pkg/provincecity/usecase"

	"github.com/gofiber/fiber/v2"
)

func ProvinceCityRoute(r fiber.Router) {
	repository := provincecityrepository.NewProvinceCityRepository()
	usecase := provincecityusecase.NewProvinceCityUseCase(repository)
	controller := provincecitycontroller.NewProvinceCityController(usecase)

	r.Get("listprovinces", controller.GetProvinces)
	r.Get("detailprovince/:province_id", controller.GetProvinceByID)
	r.Get("listcities/:province_id", controller.GetCitiesByProvinceID)
	r.Get("detailcity/:city_id", controller.GetCityByID)
}
