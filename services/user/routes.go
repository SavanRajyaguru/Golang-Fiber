package user

import (
	"github.com/gofiber/fiber/v2"
)

func UserAuthRoute(route fiber.Router) {
	auth := route.Group("/auth")

	// auth.Use(Demo)

	auth.Post("/sign-up/:id", SignUpUser)
}
