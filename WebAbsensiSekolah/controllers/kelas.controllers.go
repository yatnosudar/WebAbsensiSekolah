package controllers

import (
	"echo/WebAbsensiSekolah/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetListKelas(c echo.Context) error {
	kelas, err := models.GetListKelas()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, kelas)
}

func GetDetailKelas(c echo.Context) error {
	kelas := c.Param("kelas")

	class, err := models.GetDetailKelas(kelas)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, class)
}
