package controllers

import (
	"echo/WebAbsensiSekolah/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllGuru(c echo.Context) error {
	result, err := models.FetchAllGuru()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func StoreGuru(c echo.Context) error {
	Nama_Guru := c.FormValue("nama_guru")
	Jenis_Kelamin := c.FormValue("jenis_kelamin")
	Tanggal_Lahir := c.FormValue("tanggal_lahir")
	No_Telp := c.FormValue("no_telp")
	Kelas := c.FormValue("kelas")

	result, err := models.StoreGuru(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_Telp, Kelas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateGuru(c echo.Context) error {
	Nama_Guru := c.FormValue("nama_guru")
	Jenis_Kelamin := c.FormValue("jenis_kelamin")
	Tanggal_Lahir := c.FormValue("tanggal_lahir")
	No_Telp := c.FormValue("no_telp")
	Kelas := c.FormValue("kelas")
	Id_Guru := c.FormValue("id_guru")

	conv_Id, _ := strconv.Atoi(Id_Guru)
	result, err := models.UpdateGuru(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_Telp, Kelas, conv_Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteGuru(c echo.Context) error {
	Id_Guru := c.FormValue("id_guru")

	conv_Id, _ := strconv.Atoi(Id_Guru)

	result, err := models.DeleteGuru(conv_Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
