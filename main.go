package main

import (
	"github.com/alifnuryana/go-todo-auth/database"
	"github.com/alifnuryana/go-todo-auth/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello world! 👋🏻",
		})
	})

	router.InitRouter(app)
	database.InitDatabase()

	app.Listen(":1323")
}
