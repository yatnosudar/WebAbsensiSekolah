package controllers

import (
	"echo/WebAbsensiSekolah/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllSiswaAcc(c echo.Context) error {
	result, err := models.FetchAllSiswaAcc()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func StoreSiswaAcc(c echo.Context) error {
	Id := c.FormValue("id")
	Nis := c.FormValue("nis")
	Password := c.FormValue("password")

	conv_Id, err := strconv.Atoi(Id)
	conv_Nis, err := strconv.Atoi(Nis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreSiswaAcc(conv_Id, conv_Nis, Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSiswaAcc(c echo.Context) error {
	Nis := c.FormValue("nis")
	Password := c.FormValue("password")
	Id := c.FormValue("id")

	conv_Id, err := strconv.Atoi(Id)
	conv_Nis, err := strconv.Atoi(Nis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateSiswaAcc(conv_Nis, Password, conv_Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSiswaAcc(c echo.Context) error {
	Id := c.FormValue("id")

	conv_Id, err := strconv.Atoi(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteSiswaAcc(conv_Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
