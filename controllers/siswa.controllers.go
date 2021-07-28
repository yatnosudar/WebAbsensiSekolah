package controllers

import (
	"echo/WebAbsensiSekolah/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllSiswa(c echo.Context) error {
	result, err := models.FetchAllSiswa()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func StoreSiswa(c echo.Context) error {
	Nama_Siswa := c.FormValue("nama_siswa")
	Nis := c.FormValue("nis")
	Jenis_Kelamin := c.FormValue("jenis_kelamin")
	No_Telp := c.FormValue("no_telp")
	Kelas := c.FormValue("kelas")

	conv_Nis, err := strconv.Atoi(Nis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreSiswa(conv_Nis, Nama_Siswa, Jenis_Kelamin, No_Telp, Kelas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSiswa(c echo.Context) error {
	Nis := c.FormValue("nis")
	Jenis_Kelamin := c.FormValue("jenis_kelamin")
	No_Telp := c.FormValue("no_telp")
	Kelas := c.FormValue("kelas")
	Nama_Siswa := c.FormValue("nama_siswa")

	conv_Nis, err := strconv.Atoi(Nis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateSiswa(conv_Nis, Jenis_Kelamin, No_Telp, Kelas, Nama_Siswa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSiswa(c echo.Context) error {
	Nama_Siswa := c.FormValue("nama_siswa")

	result, err := models.DeleteSiswa(Nama_Siswa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
