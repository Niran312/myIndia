package register

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	//gorm.Model
	ID           int64     `gorm:"primaryKey;column:id" json:"-"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Mail         string    `gorm:"column:mailId;type:varchar(100);uniqueIndex;not null" json:"mail_id" validate:"required"`
	MobileNumber string    `gorm:"column:mobile_number;uniqueIndex;not null;" json:"mobile_number" validate:"required"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func Register(c *fiber.Ctx, db *gorm.DB) error {
	var user User

	validate := validator.New()

	log.Infof("User struct: ", user)

	err := c.BodyParser(&user)
	if err != nil {
		log.Errorf("Error while parsing request: ", err)
		return err
	}

	reqFields := validate.Struct(user)
	if reqFields.Error() != "nil" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"validation_error": reqFields.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("Error while hashing password: ", hashedPassword)
		return err
	}

	log.Infof("Hashed password: ", hashedPassword)

	user.Password = string(hashedPassword)

	log.Infof("User values with hashed: ", user)

	if len(user.MobileNumber) != 10 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid Mobile Number",
		})
	}

	result := db.Create(&user)

	log.Infof("Insert Query: ", result)

	if result.Error != nil {
		log.Errorf("Error while inserting into user table: ", result.Error)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"message": "User Created Successfully",
	})
}
