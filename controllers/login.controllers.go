package controllers

import (
	"echo/WebAbsensiSekolah/helper"
	"echo/WebAbsensiSekolah/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helper.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func LoginGuru(c echo.Context) error {
	// Mengisi form value dengan key
	username := c.FormValue("username")
	password := c.FormValue("password")

	response, err := models.LoginGuru(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if !response {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("guru"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": t,
	})
}

func LoginAdmin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	response, err := models.LoginAdmin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if !response {
		return echo.ErrUnauthorized
	}

	// Claim token JWT
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("admin"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	// Response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": t,
	})
}

func LoginSiswa(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Menjalankan perintah di bagian models
	response, err := models.LoginSiswa(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if !response {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("siswa"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": t,
	})
}
