package models

import (
	"database/sql"
	"echo/WebAbsensiSekolah/db"
	"fmt"
	"net/http"
)

func ListAbsenKelas(kelas string, bulan int) (Response, error) {
	var obj AbsenSiswa
	var arrobj []AbsenSiswa
	var res Response

	con := db.CreateCon()

	// menampilkan list absen berdasarkan kelas dan bulan nya
	sqlStatement := "SELECT * FROM absen WHERE id_kelas = ? AND MONTH(tanggal) = ?"

	rows, err := con.Query(sqlStatement, kelas, bulan)
	defer rows.Close()

	if err == sql.ErrNoRows {
		fmt.Println("Kelas Tidak Ditemukan")
		return res, err
	}

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Absen, &obj.Id_Guru, &obj.Nis, &obj.Id_Kelas, &obj.Absen_Masuk, &obj.Absen_Keluar, &obj.Tanggal)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}

func ListAbsenNis(Nis int, month int) (Response, error) {
	var obj AbsenSiswa
	var arrobj []AbsenSiswa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM absen WHERE nis = ? AND MONTH(tanggal) = ?"

	rows, err := con.Query(sqlStatement, Nis, month)
	defer rows.Close()

	if err == sql.ErrNoRows {
		fmt.Println("Nis Tidak Ditemukan")
		return res, err
	}

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Absen, &obj.Id_Guru, &obj.Nis, &obj.Id_Kelas, &obj.Absen_Masuk, &obj.Absen_Keluar, &obj.Tanggal)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}
