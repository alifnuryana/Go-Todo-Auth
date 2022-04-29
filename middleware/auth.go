package middleware

import (
	"errors"

	"github.com/alifnuryana/go-todo-auth/config"
	"github.com/alifnuryana/go-todo-auth/helper"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: jwtware.HS256,
		SigningKey:    []byte(config.Load("JWT_SECRET")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return helper.ResponseFailed(c, fiber.StatusBadRequest, errors.New("missing or malformed JWT"))
			}
			return helper.ResponseFailed(c, fiber.StatusUnauthorized, errors.New("invalid or expired JWT"))
		},
	})
}
