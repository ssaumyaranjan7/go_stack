package helpers

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"golang.org/x/crypto/bcrypt"
)

type userFromJWT struct {
	ID   string
	Name string
	Role string
	// Roles []string
}

func PasswordEncrypt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func ComparePassword(plainPwd string, hashedPwd []byte) bool {
	bytePlainPwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(hashedPwd, bytePlainPwd)
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetUserFromClaims(c echo.Context) *userFromJWT {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	name := claims["name"].(string)
	role := claims["role"].(string)
	return &userFromJWT{ID: id, Name: name, Role: role}
}
