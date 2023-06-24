package http

import (
	"final_project/internal/infrastructure/container"
	"final_project/internal/pkg/address"
	"final_project/internal/pkg/auth"
	"final_project/internal/pkg/category"
	"final_project/internal/pkg/store"
	"final_project/internal/pkg/user"
	"final_project/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1") // /api

	authAPI := api.Group("/auth")
	auth.AuthRoute(authAPI, containerConf)

	storeAPI := api.Group("/toko")
	store.StoreRoute(storeAPI, containerConf)

	userAPI := api.Group("/user")
	user.UserRoute(userAPI, containerConf)

	addressAPI := userAPI.Group("/alamat", utils.AuthMiddleware)
	address.AddressRoute(addressAPI, containerConf)

	categoryAPI := api.Group("/category")
	category.CategoryRoute(categoryAPI, containerConf)

}
