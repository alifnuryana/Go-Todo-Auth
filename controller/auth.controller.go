package controller

import (
	"errors"
	"strings"
	"time"

	"github.com/alifnuryana/go-todo-auth/config"
	"github.com/alifnuryana/go-todo-auth/database"
	"github.com/alifnuryana/go-todo-auth/helper"
	"github.com/alifnuryana/go-todo-auth/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func Register(c *fiber.Ctx) error {
	var requestRegister model.User
	err := c.BodyParser(&requestRegister)
	if err != nil {
		return helper.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	requestRegister.Password = helper.HashPassword(requestRegister.Password)

	tx := database.DB.Create(&requestRegister)
	if tx.Error != nil {
		return helper.ResponseFailed(c, fiber.StatusInternalServerError, tx.Error)
	}

	return helper.ResponseSuccess(c, fiber.StatusCreated, "register success")
}

func Login(c *fiber.Ctx) error {
	var requestLogin model.RequestLogin
	err := c.BodyParser(&requestLogin)
	if err != nil {
		return helper.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	var dataUser model.User
	var isValid bool
	if strings.Contains(requestLogin.Identity, "@") {
		tx := database.DB.First(&dataUser, "email = ?", requestLogin.Identity)
		isValid = helper.CheckHashPassword(requestLogin.Password, dataUser.Password)
		if tx.Error != nil {
			return helper.ResponseFailed(c, fiber.StatusUnauthorized, tx.Error)
		}
	} else {
		tx := database.DB.First(&dataUser, "username = ?", requestLogin.Identity)
		isValid = helper.CheckHashPassword(requestLogin.Password, dataUser.Password)
		if tx.Error != nil {
			return helper.ResponseFailed(c, fiber.StatusUnauthorized, tx.Error)
		}
	}

	if !isValid {
		return helper.ResponseFailed(c, fiber.StatusUnauthorized, errors.New("your credentials is not valid"))
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Username: dataUser.Username,
		Role:     dataUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Load("APP_NAME"),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	})

	token, err := tokenWithClaims.SignedString([]byte(config.Load("JWT_SECRET")))
	if err != nil {
		return helper.ResponseFailed(c, fiber.StatusInternalServerError, err)
	}

	return helper.ResponseSuccess(c, fiber.StatusOK, fiber.Map{
		"token": token,
	})
}

func Info(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return helper.ResponseSuccess(c, fiber.StatusOK, claims)
}
