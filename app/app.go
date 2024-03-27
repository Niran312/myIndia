package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"myIndia/app/routes"
	"myIndia/database"
)

func Run() {
	// Creating New Fiber Instance
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.ConfigureRoutes(app)

	database.DbInit()

	port := fmt.Sprintf(":%s", "3001")
	err := app.Listen(port)

	if err != nil {
		return
	}
}
