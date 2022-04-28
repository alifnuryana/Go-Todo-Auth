package config

import (
	"os"

	"github.com/alifnuryana/go-todo-auth/helper"
	"github.com/joho/godotenv"
)

func Load(key string) string {
	err := godotenv.Load(".env")
	helper.FatalIfError(err)
	return os.Getenv(key)
}
