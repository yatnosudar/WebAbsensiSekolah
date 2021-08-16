package models

import (
	"database/sql"
	"echo/WebAbsensiSekolah/db"
	"fmt"
	"net/http"
)

func ListAbsenKelas(kelas string) (Response, error) {
	var obj AbsenSiswa
	var arrobj []AbsenSiswa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM absen WHERE kelas = ?"

	rows, err := con.Query(sqlStatement, kelas)
	defer rows.Close()

	if err == sql.ErrNoRows {
		fmt.Println("Kelas Tidak Ditemukan")
		return res, err
	}

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Absen, &obj.Id_Guru, &obj.Nis, &obj.Kelas, &obj.Absen_Masuk, &obj.Absen_Keluar, &obj.Tanggal)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res, nil
}
