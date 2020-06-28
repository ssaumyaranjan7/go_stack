package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ssaumyaranjan7/go_stack/helpers"
	"github.com/ssaumyaranjan7/go_stack/models"
)

// CreateCab is the controller for POST:/cab/ route
func CreateCab(c echo.Context) error {
	user := helpers.GetUserFromClaims(c)
	if user.Role != "admin" {
		return c.JSON(http.StatusOK, map[string]string{
			"error": "Access Denied",
		})
	}
	cab := new(models.Cab)
	if err := c.Bind(cab); err != nil {
		return err
	}
	createdCab := models.CreateCab(cab)
	if createdCab == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "unable to create cab",
		})
	}
	return c.JSON(http.StatusOK, createdCab)
}

// GetCabDetailsByID method used for Fetching cab details using ID
func GetCabDetailsByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, map[string]string{
			"error": "cabID not found",
		})
	}
	cab := models.GetCabByID(id)
	if cab == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Cab details not found",
		})
	}
	return c.JSON(http.StatusOK, cab)
}

// FindNearByCab Gets the nearby cab
func FindNearByCab(c echo.Context) error {
	user := helpers.GetUserFromClaims(c)
	u := models.FindUserByID(user.ID)
	if u == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "User not found",
		})
	}

	cabs := models.FindNearByCab(u.Location.Longitude, u.Location.Latitude)
	return c.JSON(http.StatusOK, cabs)
}
