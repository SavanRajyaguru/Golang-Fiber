package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savan/helper"
)

func UserAuthRoute(route fiber.Router) {
	auth := route.Group("/auth")

	// auth.Use(Demo)

	auth.Post("/sign-up", helper.BodyValidator(&UserModel{}), SignUpUser)
	auth.Get("/user", GetUser)
	auth.Put("/user/:id", UpdateUser)
	auth.Delete("/user/:id", DeleteUser)

	auth.Get("/test", AssociationDemo)
}
