package models

import (
	"database/sql"
	"echo/WebAbsensiSekolah/db"
	"fmt"
)

type AbsenSiswa struct {
	Id_Absen     int    `json:"id_absen"`
	Id_Guru      int    `json:"id_guru"`
	Nis          int    `json:"nis"`
	Id_Kelas     int    `json:"id_kelas"`
	Absen_Masuk  string `json:"absen_masuk"`
	Absen_Keluar string `json:"absen_keluar"`
	Tanggal      string `json:"tanggal"`
}

func ClockIn(nis int) (bool, error) {
	var objAcc SiswaAcc

	con := db.CreateCon()

	// mendeteksi apakah nis yang diisi saat absen terdaftar/tidak
	sqlStatement := "SELECT * FROM account_siswa"

	err := con.QueryRow(sqlStatement).Scan(
		&objAcc.Id, &objAcc.Nis,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Please check your nis")
		return false, err
	}
	if err != nil {
		fmt.Println("Query error")
		return false, err
	}
	return true, nil
}

func ClockOut(nis int) (bool, error) {
	var objAcc SiswaAcc

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM account_siswa"

	err := con.QueryRow(sqlStatement).Scan(
		&objAcc.Id, &objAcc.Nis, &objAcc.Password,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Please check your nis")
		return false, err
	}
	if err != nil {
		fmt.Println("Query error")
		return false, err
	}
	return true, nil
}
