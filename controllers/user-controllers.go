package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/ssaumyaranjan7/go_stack/helpers"
	"github.com/ssaumyaranjan7/go_stack/models"
)

// Register is used from register route
func Register(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if isSucces := models.Register(u); isSucces != true {
		return c.String(http.StatusInternalServerError, "Registration fail")
	}
	return c.JSON(http.StatusOK, u)
}

// Login is used for Login route
func Login(c echo.Context) error {
	l := new(models.LoginRequest)
	if err := c.Bind(l); err != nil {
		return err
	}
	ur := models.FindUserByEmail(l.Email)
	if ur == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid username or password",
		})
	}
	fmt.Println(ur)
	check := helpers.ComparePassword(l.Password, []byte(ur.Basics.Password))
	if check != true {
		return c.String(http.StatusInternalServerError, "Login falure")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = ur.ID
	claims["name"] = ur.Basics.FirstName + " " + ur.Basics.LastName
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["role"] = ur.Role
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "login successful",
		"token":   t,
	})
}

// GetUsers is used for users route
// func GetUsers(c echo.Context) error {
// 	temp := c.Get("user")
// 	fmt.Println(temp)
// 	return c.String(http.StatusOK, "Get All User Success")
// }

// GetUserByID is used for GET: users/ route
func GetUserByID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	// roles := claims["roles"]
	// fmt.Println(roles)
	fmt.Println(id)
	u := models.FindUserByID(id)
	if u == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server error",
		})
	}
	u.Basics.Password = ""
	return c.JSON(http.StatusOK, u)
}

func UpdateUserById(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Invalid input",
		})
	}
	claims := helpers.GetUserFromClaims(c)
	u := models.FindUserByID(claims.ID)
	u.Basics.FirstName = func() string {
		if user.Basics.FirstName != "" {
			return user.Basics.FirstName
		}
		return u.Basics.FirstName

	}()
	u.Basics.LastName = func() string {
		if user.Basics.LastName != "" {
			return user.Basics.LastName
		}
		return u.Basics.LastName

	}()
	u.Basics.MobileNumber = func() string {
		if user.Basics.MobileNumber != "" {
			return user.Basics.MobileNumber
		}
		return u.Basics.MobileNumber
	}()
	u.Basics.ProfileImageURL = func() string {
		if user.Basics.ProfileImageURL != "" {
			return user.Basics.ProfileImageURL
		}
		return u.Basics.ProfileImageURL
	}()
	result := models.UpdateUser(u)
	if result == nil {
		return c.JSON(http.StatusOK, map[string]string{
			"error": "unable to update user",
		})
	}
	return c.JSON(http.StatusOK, result)
}
