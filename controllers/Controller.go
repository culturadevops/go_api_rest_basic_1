package controllers

import (
	"strconv"

	"github.com/labstack/echo"
)

type ctlHelper struct {
}

func (t *ctlHelper) Del(c echo.Context) uint {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	return uint(id)
}
