package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ssaumyaranjan7/go_stack/helpers"
	"github.com/ssaumyaranjan7/go_stack/models"
)

func CreateRide(c echo.Context) error {
	r := new(models.Ride)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "invalid input",
		})
	}
	ride := models.CreateRide(r)
	if ride == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "unable to create ride",
		})
	}
	return c.JSON(http.StatusOK, ride)
}

func GetRideByRideID(c echo.Context) error {
	rideID := c.Param("rideId")
	ride := models.GetRideByRideID(rideID)
	if ride == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Unable to find ride details",
		})
	}
	return c.JSON(http.StatusOK, ride)

}

func GetRidesByUserID(c echo.Context) error {
	user := helpers.GetUserFromClaims(c)
	fmt.Println(user.ID)
	rides := models.GetRideByUserID(user.ID)
	if rides == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "unable to find rides",
		})
	}
	return c.JSON(http.StatusOK, rides)
}
