package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/lampesm/crud-fiber/db"
	_ "github.com/lampesm/crud-fiber/docs"
	"github.com/lampesm/crud-fiber/serializers"
)

// @title CRUD Fiber
// @version 1.0
// @description This is a  swagger for CURD Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:3005
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/accounts/:id", ShowAccount)
	app.Post("/user/create", ShowAccount)

	app.Listen(":3005")
}

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

func CreateUser(c *fiber.Ctx) error {
	//payload := new(serializers.RegisterRequest)
}
