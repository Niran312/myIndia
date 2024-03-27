package state

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type s struct {
	gorm.Model
	AllStates []states `json:"States"`
}
type states struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Capital     string `json:"capital"`
	Population  string `json:"population"`
	RulingParty string `json:"ruling_party"`
	Website     string `json:"official_website"`
	Description string `json:"description"`
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
