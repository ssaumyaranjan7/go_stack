package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ssaumyaranjan7/go_stack/controllers"
)

func RideSubRoutes(g *echo.Group) {
	g.Use(middleware.JWT([]byte("secret")))
	g.POST("/", controllers.CreateRide)
	g.GET("/", controllers.GetRidesByUserID)
	g.GET("/:id", controllers.GetRideByRideID)
}
