package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"myIndia/app/routes"
	cfg "myIndia/configuration"
	"myIndia/database"
)

func Run() {
	app := fiber.New()

	cfg.LoadConfig()
	config := cfg.GetConfig()

	db := database.DbInit()
	//log.Printf("Db values: %v", db)

	//sdf := db.AutoMigrate(&register.User{})

	migrateDb := db.AutoMigrate(&routes.User{})
	//if migrateDb.Error() != "nil" {
	//	log.Printf("Error while auto migrate: %v", migrateDb.Error())
	//	return
	//}
	fmt.Printf("Migrate db: %v", migrateDb)

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.ConfigureRoutes(app, db)

	//database.DbInit()

	port := fmt.Sprintf(":%s", config.Port)
	err := app.Listen(port)

	if err != nil {
		return
	}
}
