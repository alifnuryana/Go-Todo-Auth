package controller

import (
	"github.com/alifnuryana/go-todo-auth/database"
	"github.com/alifnuryana/go-todo-auth/helper"
	"github.com/alifnuryana/go-todo-auth/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []model.Todo
	tx := database.DB.Find(&todos)
	if tx.Error != nil {
		return helper.ResponseFailed(c, fiber.StatusNotFound, tx.Error)
	}
	return helper.ResponseSuccess(c, fiber.StatusOK, todos)
}

func GetTodo(c *fiber.Ctx) error {
	todoId, err := c.ParamsInt("todoId")
	if err != nil {
		return helper.ResponseFailed(c, fiber.StatusInternalServerError, err)
	}

	var todo model.Todo
	tx := database.DB.First(&todo, todoId)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return helper.ResponseSuccess(c, fiber.StatusOK, gorm.ErrRecordNotFound.Error())
		}
		return helper.ResponseFailed(c, fiber.StatusInternalServerError, tx.Error)
	}

	return helper.ResponseSuccess(c, fiber.StatusOK, todo)
}

func PostTodo(c *fiber.Ctx) error {
	var todo model.Todo
	err := c.BodyParser(&todo)
	if err != nil {
		return helper.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	tx := database.DB.Create(&todo)
	if tx.Error != nil {
		return helper.ResponseFailed(c, fiber.StatusInternalServerError, tx.Error)
	}

	return helper.ResponseSuccess(c, fiber.StatusOK, todo)
}

func PutTodo(c *fiber.Ctx) error {
	todoId, err := c.ParamsInt("todoId")
	if err != nil {
		return helper.ResponseFailed(c, fiber.StatusInternalServerError, err)
	}

	var todo model.Todo
	err = c.BodyParser(&todo)
	if err != nil {
		return helper.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	tx := database.DB.Model(&model.Todo{}).Where("id = ?", todoId).Updates(todo)
	if tx.Error != nil {
		return helper.ResponseFailed(c, fiber.StatusInternalServerError, tx.Error)
	}

	return helper.ResponseSuccess(c, fiber.StatusOK, todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	todoId, err := c.ParamsInt("todoId")
	if err != nil {
		return helper.ResponseFailed(c, fiber.StatusInternalServerError, err)
	}

	tx := database.DB.Delete(&model.Todo{}, todoId)
	if tx.Error != nil {
		return helper.ResponseFailed(c, fiber.StatusInternalServerError, tx.Error)
	}

	return helper.ResponseSuccess(c, fiber.StatusOK, "successfully deleted.")
}
