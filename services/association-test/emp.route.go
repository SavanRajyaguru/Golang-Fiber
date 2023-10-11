package associationtest

import (
	"github.com/gofiber/fiber/v2"
)

func EmpTestRoute(route fiber.Router) {
	auth := route.Group("/emp")

	auth.Get("/", GetEmp)
	auth.Post("/", TransactionDemo)
}
