package controllers

import (
	"net/http"
	"strconv"

	"github.com/culturadevops/api/models"
	"github.com/labstack/echo"
)

var User *userController

type userController struct {
}

type RequestUser struct {
	Account  string `json:"account" form:"account" query:"account"`
	Password string `json:"password" form:"password" query:"password"`
}

func (t *userController) Add(c echo.Context) error {
	u := new(RequestUser)
	if err := c.Bind(u); err != nil {
		return err
	}
	models.User.Add(u.Account, u.Password)
	return c.JSON(http.StatusOK, "Registro Creado")
}
func (t *userController) List(c echo.Context) error {
	list := models.User.List()
	return c.JSON(http.StatusOK, list)
}
func (t *userController) Show(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	use, _ := models.User.Info(uint(id))
	return c.JSON(http.StatusOK, use)
}

func (t *userController) Del(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	models.User.Del(uint(id))
	return c.JSON(http.StatusOK, "Registro eliminado")
}
func (t *userController) Update(c echo.Context) error {

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	u := new(RequestUser)
	if err := c.Bind(u); err != nil {
		return err
	}
	models.User.Update(uint(id), u.Account, u.Password)

	return c.JSON(http.StatusOK, "Registro actualizado")
}
