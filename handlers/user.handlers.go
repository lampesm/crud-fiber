package handlers

import (
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lampesm/crud-fiber/db"
	"github.com/lampesm/crud-fiber/entity"
	"github.com/lampesm/crud-fiber/serializers"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @TAGS account
// @ID Show-account
// @Accept  json
// @Produce  json
// @Param id path string true "User.ID"
// @Success 200
// @Router /api/v1/account/read/{id} [get]
func ShowAccount(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	_ = claims

	posDB := db.Connection()
	defer db.Close(posDB)

	var result Result
	posDB.Table("users").Select("id", "username", "password", "email").Where("id = ?", c.Params("id")).Scan(&result)

	if result.ID != 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"msg":   result,
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": false,
			"msg":   "user not found",
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
// @TAGS account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param payload body serializers.User true "User"
// @Success 200
// @Router /api/v1/account/create [post]
func CreateAcount(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	_ = claims

	payload := new(serializers.User)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
			"msg":   errDB.Error.Error(),
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error": false,
			"msg":   "inserted",
		})
	}
}

// UpdateUser godoc
// @Summary update a account
// @Description update a account
// @TAGS account
// @ID update-account
// @Accept  json
// @Produce  json
// @Param payload body serializers.User true "User"
// @Param id path string true "User.ID"
// @Success 200
// @Router /api/v1/account/update/{id} [put]
func UpdateAccount(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	_ = claims

	payload := new(serializers.User)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	posDB := db.Connection()
	defer db.Close(posDB)

	_, errEmail := mail.ParseAddress(payload.Email)
	if errEmail != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "email is not valid",
		})
	}

	result := posDB.Table("users").Where("id = ?", c.Params("id")).Updates(map[string]interface{}{
		"username": payload.Username, "password": payload.Password, "email": payload.Email,
	})

	if result.RowsAffected == 1 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"msg":   "user updated",
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": result.Error.Error(),
			"msg":   "user not found",
		})
	}
}

// DeleteUser godoc
// @Summary delete a account
// @Description delete a account
// @TAGS account
// @ID delete-account
// @Accept  json
// @Produce  json
// @Param id path string true "User.ID"
// @Success 200
// @Router /api/v1/account/delete/{id} [delete]
func DeleteAccount(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	_ = claims
	
	var users []entity.User
	posDB := db.Connection()
	defer db.Close(posDB)

	result := posDB.Where("id = ?", c.Params("id")).Delete(&users)

	if result.RowsAffected == 1 {
		return c.Status(fiber.StatusNoContent).Send(nil)
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user not found",
		})
	}
}
