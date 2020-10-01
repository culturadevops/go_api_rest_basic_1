package routes

import (
	"github.com/culturadevops/api/controllers"
	"github.com/labstack/echo"
)

// ejemplo para usar en main.go	routes.User(e)
func User(e *echo.Echo) {
	r := e.Group("/user")
	/*middlewares*/
	r.POST("/add", controllers.User.Add)
	r.PUT("/update/:id", controllers.User.Update)
	r.DELETE("/del/:id", controllers.User.Del)
	r.GET("/show/:id", controllers.User.Show)
	r.GET("/list", controllers.User.List)
}
