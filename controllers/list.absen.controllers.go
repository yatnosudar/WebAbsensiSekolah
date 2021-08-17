package controllers

import (
	"echo/WebAbsensiSekolah/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ListAbsenKelas(c echo.Context) error {
	kelas := c.Param("kelas")
	bulan := c.Param("bulan")

	conv_bulan, _ := strconv.Atoi(bulan)

	class, err := models.ListAbsenKelas(kelas, conv_bulan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, class)
}
