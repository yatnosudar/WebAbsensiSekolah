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

func FetchAllGuru(c echo.Context) error {
	guru := models.GuruDetail{}

	res := []models.GuruDetail{}
	db := db.CreateCon()

	sqlStatement := "SELECT id_guru, nama_guru, jenis_kelamin,tanggal_lahir, no_telp FROM guru"

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage := 2

	var total int64
	db.QueryRow("SELECT count(id_guru) FROM guru ORDER by id_guru").Scan(&total)

	sqlStatement = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlStatement, perPage, (page)*perPage)

	GuruRows, err := db.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	for GuruRows.Next() {
		err = GuruRows.Scan(&guru.Id_Guru, &guru.Nama_Guru, &guru.Jenis_Kelamin, &guru.Tanggal_Lahir, &guru.No_Telp)

		res = append(res, guru)
	}

	db.QueryRow(sqlStatement).Scan(&res)

	response := make(map[string]interface{}, 4)
	response["data"] = res
	response["total_data"] = total
	response["page"] = page
	response["last_page"] = math.Ceil(float64(total / int64(perPage)))

	return c.JSON(http.StatusOK, response)
}

func StoreGuru(c echo.Context) error {
	Nama_Guru := c.FormValue("nama_guru")
	Jenis_Kelamin := c.FormValue("jenis_kelamin")
	Tanggal_Lahir := c.FormValue("tanggal_lahir")
	No_Telp := c.FormValue("no_telp")
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	Role := c.FormValue("role")

	result, err := models.StoreGuru(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_Telp, Username, Password, Role)
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
	Username := c.FormValue("username")
	Password := c.FormValue("password")
	Role := c.FormValue("role")
	Id_Guru := c.FormValue("id_guru")

	conv_Id, _ := strconv.Atoi(Id_Guru)
	result, err := models.UpdateGuru(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_Telp, Username, Password, Role, conv_Id)
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
