package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ssaumyaranjan7/go_stack/controllers"
)

func CabSubRoutes(g *echo.Group) {
	g.Use(middleware.JWT([]byte("secret")))
	g.POST("/", controllers.CreateCab)
	g.GET("/nearByCabs", controllers.FindNearByCab)
	g.GET("/:id", controllers.GetCabDetailsByID)
}
