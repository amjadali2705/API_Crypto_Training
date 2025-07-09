package router

import (
	"gormapp/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/users")
	api.Get("/", controller.GetUsers)
	api.Get("/:id", controller.GetUser)
	api.Post("/", controller.CreateUser)
	api.Put("/:id", controller.UpdateUser)
	api.Delete("/:id", controller.DeleteUser)

	profiles := app.Group("/profiles")
	profiles.Get("/", controller.GetProfiles)
	profiles.Get("/:id", controller.GetProfile)
	profiles.Post("/", controller.CreateProfile)
	profiles.Put("/:id", controller.UpdateProfile)
	profiles.Delete("/:id", controller.DeleteProfile)

	orders := app.Group("/orders")
	orders.Get("/", controller.GetOrders)
	orders.Get("/:id", controller.GetOrder)
	orders.Post("/", controller.CreateOrder)
	orders.Put("/:id", controller.UpdateOrder)
	orders.Delete("/:id", controller.DeleteOrder)
}
