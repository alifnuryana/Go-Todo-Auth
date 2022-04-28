package router

import (
	"github.com/alifnuryana/go-todo-auth/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	todos := api.Group("/todos")
	todos.Get("/", controller.GetTodos)
	todos.Get("/:todoId", controller.GetTodo)
	todos.Post("/", controller.PostTodo)
	todos.Put("/:todoId", controller.PutTodo)
	todos.Delete("/:todoId", controller.DeleteTodo)
}
