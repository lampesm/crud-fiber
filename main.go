package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lampesm/crud-fiber/docs"
	"github.com/lampesm/crud-fiber/handlers"
)

// @title CRUD Fiber
// @version 1.0
// @description This is a  swagger for CURD Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:3005/api/v1
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	api_v1 := app.Group("/api/v1")

	api_v1.Get("/account/read/:id", handlers.ShowAccount)
	api_v1.Post("/account/create", handlers.CreateAcount)
	api_v1.Put("/account/update/:id", handlers.UpdateAccount)
	api_v1.Delete("/account/delete/:id", handlers.DeleteAccount)

	app.Listen(":3005")
}
