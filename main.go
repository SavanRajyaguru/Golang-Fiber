package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/savan/config"
	"github.com/savan/database"
	"github.com/savan/middleware"
	"github.com/savan/migration"
)

func init() {
	config.InitEnvVariables()
	database.ConnectDB()
	migration.LoadAllSchema()
}

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(middleware.RecoveryMiddleware)

	middleware.RegisterRoutes(app)
	// For the hot reload script
	// nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run ./main.go
	log.Println("Server running on...", config.ConfigEnv.PORT)
	log.Fatal(app.Listen(":" + config.ConfigEnv.PORT).Error())
}
