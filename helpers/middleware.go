package helpers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func JWTAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		temp := c.Get("user")
		fmt.Println(temp)
		user := c.Get("user").(*jwt.Token)
		fmt.Println(user)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		ID := claims["ID"]
		fmt.Println(name)
		fmt.Println(ID)
		next(c)
		return nil
	}
}
