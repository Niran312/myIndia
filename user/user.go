package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"myIndia/register"
)

func DataUser(c *fiber.Ctx, db *gorm.DB) error {

	var user register.User

	userID := c.Params("id")

	result := db.Find(&user, "id=?", userID)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"data": user,
	})
}
