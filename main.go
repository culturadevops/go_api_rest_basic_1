package main

import (
	"github.com/culturadevops/api/libs"
	"github.com/culturadevops/api/routes"
	"github.com/labstack/echo"
	config "github.com/spf13/viper"
)

func init() {
	config.AddConfigPath("./configs")
	dbConfig := libs.Configure("./configs", "mysql")
	libs.DB = dbConfig.InitMysqlDB()
}
func mainMiddlewares(e *echo.Echo) {

}
func mainRoutes(e *echo.Echo) {
	routes.User(e)
}
func main() {
	e := echo.New()
	mainMiddlewares(e)
	mainRoutes(e)
	e.Logger.Fatal(e.Start(":80"))
}
