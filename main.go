package main

import (
	"fmt"

	"github.com/Kamva/mgm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ssaumyaranjan7/go_stack/controllers"
	"github.com/ssaumyaranjan7/go_stack/routes"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	// register Users Routes
	routes.UserSubRoutes(e.Group("/users"))
	routes.CabSubRoutes(e.Group("/cab"))
	routes.RideSubRoutes(e.Group("/ride"))
	e.POST("/login", controllers.Login)
	e.POST("/users/", controllers.Register)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func init() {
	// Setup mgm default config
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://localhost:27017/stack"))
	if err != nil {
		fmt.Println("connected")
	}
}
