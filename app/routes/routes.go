package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	reg "myIndia/register"
	"myIndia/user"
	"time"
)

type User struct {
	//gorm.Model
	ID           int64     `gorm:"primaryKey;column:id" json:"-"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Mail         string    `gorm:"column:mailId;type:varchar(100);uniqueIndex;not null" json:"mail_id"`
	MobileNumber string    `gorm:"column:mobile_number;uniqueIndex;not null;" json:"mobile_number"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func AllStates(c *fiber.Ctx) error {
	err := c.SendString("States will be shown here")

	return err
}

func AllDistricts(c *fiber.Ctx) error {
	err := c.SendString("District will be shown here")

	return err
}

func AllCities(c *fiber.Ctx) error {
	err := c.SendString("Cities will be shown here")

	return err
}

func ConfigureRoutes(app *fiber.App, db *gorm.DB) {

	app.Post("/register", func(c *fiber.Ctx) error {
		return reg.Register(c, db)
	})
	app.Get("/user/:id", func(ctx *fiber.Ctx) error {
		return user.DataUser(ctx, db)
	})
	app.Get("/all_states", AllStates)
	app.Get("/all_districts", AllDistricts)
	app.Get("/all_cities", AllCities)
}
