package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(passowrd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passowrd), 14)
	return string(bytes), err
}

func CheckPasswordHash(passowrd string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passowrd))
	return err == nil
}
