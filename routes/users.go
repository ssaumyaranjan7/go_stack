package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ssaumyaranjan7/go_stack/controllers"
)

// UserSubRoutes is for Sub route grouping
func UserSubRoutes(g *echo.Group) {
	// g.Use(helpers.JWTAuthentication)
	g.Use(middleware.JWT([]byte("secret")))
	// g.GET("/", controllers.GetUsers)
	g.GET("/", controllers.GetUserByID)
	// g.POST("/login", controllers.Login)

	g.PUT("/", controllers.UpdateUserById)

}
