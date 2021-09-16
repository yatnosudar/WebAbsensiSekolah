package controllers

import (
	"echo/WebAbsensiSekolah/models"
	"net/http"
	"strconv"

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
	// menggunakan parameter untuk memilih kelas
	kelas := c.Param("kelas")

	class, err := models.GetDetailKelas(kelas)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, class)
}

func AddKelas(c echo.Context) error {
	Kelas := c.FormValue("kelas")
	Id_guru := c.FormValue("id_guru")

	// konversi dari string ke int
	conv_Id_Guru, _ := strconv.Atoi(Id_guru)

	// mengisi parameter di package models dengan fom value di atas
	result, err := models.AddKelas(Kelas, conv_Id_Guru)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateKelas(c echo.Context) error {
	Kelas := c.FormValue("kelas")
	Id_Guru := c.FormValue("id_guru")
	Id_Kelas := c.FormValue("id_kelas")

	conv_Id_Guru, _ := strconv.Atoi(Id_Guru)
	conv_Id_Kelas, _ := strconv.Atoi(Id_Kelas)

	result, err := models.UpdateKelas(Kelas, conv_Id_Guru, conv_Id_Kelas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteKelas(c echo.Context) error {
	Id_Kelas := c.FormValue("id_kelas")

	conv_Id, _ := strconv.Atoi(Id_Kelas)

	result, err := models.DeleteKelas(conv_Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
