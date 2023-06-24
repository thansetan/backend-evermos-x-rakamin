package address

import (
	"final_project/internal/infrastructure/container"
	addresscontroller "final_project/internal/pkg/address/controller"
	addressrepository "final_project/internal/pkg/address/repository"
	addressusecase "final_project/internal/pkg/address/usecase"

	"github.com/gofiber/fiber/v2"
)

func AddressRoute(r fiber.Router, conf *container.Container) {
	repository := addressrepository.NewAddressRepository(conf.Mysqldb)
	usecase := addressusecase.NewAddressUseCase(repository)
	controller := addresscontroller.NewAddressController(usecase)

	r.Get("", controller.GetMyAddresses)
	r.Get(":address_id", controller.GetMyAddressByID)
	r.Post("", controller.CreateAddress)
	r.Put(":address_id", controller.UpdateMyAddressByID)
	r.Delete(":address_id", controller.DeleteMyAddressByID)
}
