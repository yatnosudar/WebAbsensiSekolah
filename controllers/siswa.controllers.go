package controllers

import (
	"echo/WebAbsensiSekolah/db"
	"echo/WebAbsensiSekolah/models"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllSiswa(c echo.Context) error {
	siswa := models.SiswaDetail{}

	res := []models.SiswaDetail{}
	db := db.CreateCon()

	sqlStatement := "SELECT * FROM siswa"

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage := 2

	var total int64
	db.QueryRow("SELECT count(nis) FROM siswa ORDER by nis").Scan(&total)

	sqlStatement = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlStatement, perPage, (page)*perPage)

	SiswaRows, err := db.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	for SiswaRows.Next() {
		err = SiswaRows.Scan(&siswa.Id_Siswa, &siswa.Nama_Siswa, &siswa.Nis, &siswa.Jenis_Kelamin, &siswa.No_Telp, &siswa.Kelas)

		res = append(res, siswa)
	}

	db.QueryRow(sqlStatement).Scan(&res)

	response := make(map[string]interface{}, 4)
	response["data"] = res
	response["total_data"] = total
	response["page"] = page
	response["last_page"] = math.Ceil(float64(total / int64(perPage)))

	return c.JSON(http.StatusOK, response)
}

func AddSiswa(c echo.Context) error {
	Nama_Siswa := c.FormValue("nama_siswa")
	Nis := c.FormValue("nis")
	Jenis_Kelamin := c.FormValue("jenis_kelamin")
	No_Telp := c.FormValue("no_telp")
	Kelas := c.FormValue("kelas")

	conv_Nis, err := strconv.Atoi(Nis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.AddSiswa(Nama_Siswa, conv_Nis, Jenis_Kelamin, No_Telp, Kelas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSiswa(c echo.Context) error {
	Nama_Siswa := c.FormValue("nama_siswa")
	Nis := c.FormValue("nis")
	Jenis_Kelamin := c.FormValue("jenis_kelamin")
	No_Telp := c.FormValue("no_telp")
	Kelas := c.FormValue("kelas")
	Id_Siswa := c.FormValue("id_siswa")

	conv_Id, err := strconv.Atoi(Id_Siswa)
	conv_Nis, err := strconv.Atoi(Nis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// mengisi parameter dari package models dengan form value
	result, err := models.UpdateSiswa(Nama_Siswa, conv_Nis, Jenis_Kelamin, No_Telp, Kelas, conv_Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSiswa(c echo.Context) error {
	Id_Siswa := c.FormValue("id_siswa")

	conv_Id, _ := strconv.Atoi(Id_Siswa)

	result, err := models.DeleteSiswa(conv_Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
