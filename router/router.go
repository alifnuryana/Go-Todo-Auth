package router

import (
	"github.com/alifnuryana/go-todo-auth/controller"
	"github.com/alifnuryana/go-todo-auth/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	todos := api.Group("/todos")
	todos.Get("/", middleware.Protected(), controller.GetTodos)
	todos.Get("/:todoId", middleware.Protected(), controller.GetTodo)
	todos.Post("/", middleware.Protected(), controller.PostTodo)
	todos.Put("/:todoId", middleware.Protected(), controller.PutTodo)
	todos.Delete("/:todoId", middleware.Protected(), controller.DeleteTodo)

	auth := api.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Get("/info", middleware.Protected(), controller.Info)
}
