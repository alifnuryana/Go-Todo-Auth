package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	byteString, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	FatalIfError(err)
	return string(byteString)
}

func CheckHashPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
