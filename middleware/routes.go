package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/savan/helper"
	associationtest "github.com/savan/services/association-test"
	"github.com/savan/services/jwt"
	"github.com/savan/services/user/auth"
)

func registerV1Routes(apiRoute fiber.Router) {
	v1Route := apiRoute.Group("/v1")

	auth.UserAuthRoute(v1Route)
	associationtest.EmpTestRoute(v1Route)
	jwt.JwtTestRoute(v1Route)
}

func RegisterRoutes(app *fiber.App) {
	apiRoute := app.Group("/api")

	registerV1Routes(apiRoute)

	app.All("*", func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(helper.Response{
			Code:    fiber.StatusNotFound,
			Message: "Route Not Found!!",
		})
	})
}
