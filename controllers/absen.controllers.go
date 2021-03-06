package controllers

import (
	"echo/WebAbsensiSekolah/db"
	"echo/WebAbsensiSekolah/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type AbsenSiswa struct {
	Id_Absen     int    `json:"id_absen"`
	Id_Guru      int    `json:"id_guru"`
	Nis          int    `json:"nis"`
	Id_Kelas     string `json:"id_kelas"`
	Absen_Masuk  string `json:"absen_masuk"`
	Absen_Keluar string `json:"absen_keluar"`
	Tanggal      string `json:"tanggal"`
}

func ClockIn(c echo.Context) error {
	con := db.CreateCon()

	nis := c.FormValue("nis")

	conv_nis, _ := strconv.Atoi(nis)

	response, err := models.ClockIn(conv_nis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if !response {
		return echo.ErrUnauthorized
	}

	var obj AbsenSiswa

	// perintah sql untuk mendapatkan kelas dari nis yang diisi tadi
	sqlStatementKelas := "SELECT kelas FROM siswa WHERE nis = ?"

	errKelas := con.QueryRow(sqlStatementKelas, &conv_nis).Scan(
		&obj.Id_Kelas,
	)

	if errKelas != nil {
		fmt.Println("Query error Kelas")
		return errKelas
	}

	// perintah sql untuk mendapatkan id guru yang menjadi wali dari kelas di atas
	sqlStatementIdGuru := "SELECT id_guru FROM kelas WHERE id_kelas = ?"

	errGuru := con.QueryRow(sqlStatementIdGuru, &obj.Id_Kelas).Scan(
		&obj.Id_Guru,
	)

	if errGuru != nil {
		fmt.Println("Query error Id")
		return errGuru
	}

	// untuk mendapatkan waktu hari
	waktu := time.Now()

	// untuk mendapatkan data waktu sekarang sesuai format mysql
	waktu_sekarang := waktu.Format(time.RFC3339)

	// perintah sql untuk memasukkan data ke dalam tabel absen saat siswa melakukan absen
	sqlStatementAbsen := "INSERT absen (id_guru, nis, kelas, absen_masuk, tanggal) VALUES (?, ?, ? , ?, ?)"

	stmt, err := con.Prepare(sqlStatementAbsen)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(obj.Id_Guru, conv_nis, obj.Id_Kelas, waktu_sekarang, waktu.Format("2006-01-02"))
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()

	fmt.Sprintf("Last inserted id : %s", id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Anda telah berhasil melakukan absen masuk pada tanggal ": waktu.Format("2006-01-02"),
	})
}

func ClockOut(c echo.Context) error {
	con := db.CreateCon()

	nis := c.FormValue("nis")

	conv_nis, _ := strconv.Atoi(nis)

	response, err := models.ClockOut(conv_nis)
	if err != nil {
		return err
	}
	if !response {
		return echo.ErrUnauthorized
	}

	// update data absen berdasarkan per hari ini dan menentukan baris nya dengan nis yang diisi
	sqlStatementAbsen := "UPDATE absen SET absen_keluar = ? WHERE nis = ? AND tanggal = ?"

	stmt, err := con.Prepare(sqlStatementAbsen)
	if err != nil {
		return err
	}

	waktu := time.Now()
	waktu_sekarang := waktu.Format(time.RFC3339)

	result, err := stmt.Exec(waktu_sekarang, conv_nis, waktu.Format("2006-01-02"))
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Sprintf("RowsAffected : %d", rowsAffected)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Anda telah berhasil melakukan absen keluar pada tanggal": waktu.Format("2006-01-02"),
	})
}
