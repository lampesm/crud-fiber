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
// @Param id path int true "Account ID"
// @Success 200 {object} Account
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /accounts/{id} [get]
func ShowAccount(c *fiber.Ctx) error {
	return c.JSON(Account{
		Id: c.Params("id"),
	})
}

type Account struct {
	Id string
}

type HTTPError struct {
	status  string
	message string
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
