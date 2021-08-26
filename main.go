package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/laksharzky/golang-fiber-basic-todo-app/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	//setup routes
	setupRoutes(app)

	//listen on server
	err := app.Listen(":8000")

	//handle error
	if err != nil {
		panic(err)
	}

}

func setupRoutes(app *fiber.App) {
	//give respons when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	//api group
	api := app.Group("/api")

	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint",
		})
	})

	//connect todo routes
	routes.TodoRoute(api.Group("/todos"))
}
