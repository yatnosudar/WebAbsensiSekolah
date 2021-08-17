package routes

import (
	"echo/WebAbsensiSekolah/controllers"
	"echo/WebAbsensiSekolah/middlewares"

	// "echo/WebAbsensiSekolah/middlewares"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	// Kelas
	e.GET("/kelas", controllers.GetListKelas)
	e.GET("/kelas/:kelas", controllers.GetDetailKelas)
	e.POST("/kelas/add", controllers.AddKelas, middlewares.LoginAdmin)
	e.PUT("/kelas/edit", controllers.UpdateKelas, middlewares.LoginAdmin)
	e.DELETE("/kelas/delete", controllers.DeleteKelas, middlewares.LoginAdmin)

	// Guru
	e.GET("/guru", controllers.FetchAllGuru)
	e.POST("/guru/", controllers.StoreGuru)
	e.PUT("/guru/", controllers.UpdateGuru)
	e.DELETE("/guru/", controllers.DeleteGuru)

	// Siswa
	e.GET("/siswa", controllers.FetchAllSiswa)
	e.POST("/siswa/", controllers.StoreSiswa)
	e.PUT("/siswa/", controllers.UpdateSiswa)
	e.DELETE("/siswa/", controllers.DeleteSiswa, middlewares.LoginAdmin)

	// Acc Siswa
	e.GET("/accsiswa", controllers.FetchAllSiswaAcc, middlewares.LoginAdmin)
	e.POST("/accsiswa/", controllers.StoreSiswaAcc, middlewares.LoginAdmin)
	e.PUT("/accsiswa/", controllers.UpdateSiswaAcc)
	e.DELETE("/accsiswa/", controllers.DeleteSiswaAcc, middlewares.LoginAdmin)

	// Absen Siswa
	e.POST("/absen/clock-in", controllers.ClockIn, middlewares.LoginSiswa)
	e.POST("/absen/clock-out", controllers.ClockOut, middlewares.LoginSiswa)
	e.GET("/absen/list-absen/:kelas/:bulan", controllers.ListAbsenKelas, middlewares.LoginGuru)

	// Membuat hash password
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)

	// Login
	e.POST("/login-guru", controllers.LoginGuru)
	e.POST("/login-siswa", controllers.LoginSiswa)
	e.POST("/login-admin", controllers.LoginAdmin)

	return e
}
