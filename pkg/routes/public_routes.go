package routes

import (
	"go-fiber-tutor/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/user", controllers.UserSignUp)
	route.Get("/user", controllers.GetAllUser)
	route.Put("/user/:id", controllers.UpdateUser)
	route.Delete("/user/:id", controllers.DeleteUser)
	route.Post("/user/login", controllers.UserSignIn)
}
