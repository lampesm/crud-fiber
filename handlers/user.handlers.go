package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lampesm/crud-fiber/db"
	"github.com/lampesm/crud-fiber/serializers"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path string true "User.ID"
// @Success 200
// @Router /account/{id} [get]
func ShowAccount(c *fiber.Ctx) error {
	posDB := db.Connection()
	defer db.Close(posDB)

	var result Result
	posDB.Table("users").Select("id", "username", "password", "email").Where("id = ?", c.Params("id")).Scan(&result)

	if result.ID != 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"content": result,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"content": "user not found",
		})
	}
}

type Result struct {
	ID       int
	Username string
	Password string
	Email    string
}

// CreateUser godoc
// @Summary create a account
// @Description create a account
// @Accept  json
// @Produce  json
// @Param payload body serializers.User true "User"
// @Success 200
// @Router /user/create [post]
func CreateUser(c *fiber.Ctx) error {
	payload := new(serializers.User)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	posDB := db.Connection()
	defer db.Close(posDB)

	errDB := posDB.Create(&payload)
	if errDB.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   errDB.Error,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"msg":   "inserted",
		})
	}
}
