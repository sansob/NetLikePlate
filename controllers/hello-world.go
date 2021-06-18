package controllers

import (
	"NetLikePlate/properties"
	"github.com/labstack/echo/v4"
	"net/http"
)

type response1 struct {
	Page   int
	Fruits []string
}

func HelloWorld(c echo.Context) error {
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	return c.JSON(http.StatusOK, properties.NResponse("OK", "A11", res1D))
}

func Testing123(bil1 int, bil2 int) int {
	var x = bil1 + bil2
	return x
}
