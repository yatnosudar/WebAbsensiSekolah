package controllers

import (
	"echo/WebAbsensiSekolah/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListAbsenKelas(c echo.Context) error {
	kelas := c.Param("kelas")

	class, err := models.ListAbsenKelas(kelas)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, class)
}
