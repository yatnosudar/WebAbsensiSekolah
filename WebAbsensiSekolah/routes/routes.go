package routes

import (
	"echo/WebAbsensiSekolah/controllers"
	"echo/WebAbsensiSekolah/middlewares"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	// Kelas
	e.GET("/kelas", controllers.GetListKelas)
	e.GET("/kelas/:kelas", controllers.GetDetailKelas)

	// Guru
	e.GET("/guru", controllers.FetchAllGuru)
	e.POST("/guru/", controllers.StoreGuru)
	e.PUT("/guru/", controllers.UpdateGuru)
	e.DELETE("/guru/", controllers.DeleteGuru)

	// Siswa
	e.GET("/siswa", controllers.FetchAllSiswa)
	e.POST("/siswa/", controllers.StoreSiswa)
	e.PUT("/siswa/", controllers.UpdateSiswa)
	e.DELETE("/siswa/", controllers.DeleteSiswa)

	// Acc Siswa
	e.GET("/accsiswa", controllers.FetchAllSiswaAcc, middlewares.LoginAdmin)
	e.POST("/accsiswa/", controllers.StoreSiswaAcc)
	e.PUT("/accsiswa/", controllers.UpdateSiswaAcc)
	e.DELETE("/accsiswa/", controllers.DeleteSiswaAcc)

	// Login
	e.POST("/login-admin", controllers.CheckLoginAdmin)
	e.POST("/login-siswa", controllers.CheckLoginSiswa)

	return e
}
