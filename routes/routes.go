package routes

import (
	"echo/WebAbsensiSekolah/controllers"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	// Kelas
	e.GET("/kelas", controllers.GetListKelas)
	e.GET("/kelas/:kelas", controllers.GetDetailKelas)

	return e
}
