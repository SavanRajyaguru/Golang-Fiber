package jwt

import (
	"github.com/gofiber/fiber/v2"
)

func JwtTestRoute(route fiber.Router) {
	routes := route.Group("/jwt")

	routes.Get("/", GenerateJwt)
}
