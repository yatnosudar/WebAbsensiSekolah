package models

import (
	"database/sql"
	"echo/WebAbsensiSekolah/db"
	"fmt"
	"net/http"
)

type Kelas struct {
	Id_Kelas int    `json:"id_kelas"`
	Kelas    string `json:"kelas"`
}

type Siswa struct {
	Id_Siswa      int    `json:"id_siswa"`
	Nama_Siswa    string `json:"nama_siswa"`
	Nis           int    `json:"nis"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	No_Telp       string `json:"no_telp"`
	Kelas         string `json:"kelas"`
}

type Guru struct {
	Id_Guru       int    `json:"id_guru"`
	Nama_Guru     string `json:"nama_guru"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Tanggal_Lahir string `json:"tanggal_lahir"`
	No_Telp       string `json:"no_telp"`
	Kelas         string `json:"kelas"`
}

func GetListKelas() (Response, error) {
	var obj Kelas
	var arrobj []Kelas
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM kelas"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Kelas, &obj.Kelas)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res, nil
}

func GetDetailKelas(kelas string) (Response, error) {
	var objKelas Kelas

	var objGuru Guru

	var objSiswa Siswa
	var arrobjSiswa []Siswa
	var res Response

	con := db.CreateCon()

	sqlStatementKelas := "SELECT * FROM kelas WHERE kelas =?"

	errKelas := con.QueryRow(sqlStatementKelas, kelas).Scan(
		&objKelas.Id_Kelas, &objKelas.Kelas,
	)

	if errKelas == sql.ErrNoRows {
		fmt.Println("Kelas Tidak Ditemukan")
		return res, errKelas
	}

	if errKelas != nil {
		return res, errKelas
	}

	sqlStatementGuru := "SELECT * FROM guru WHERE kelas =?"

	errGuru := con.QueryRow(sqlStatementGuru, kelas).Scan(
		&objGuru.Id_Guru, &objGuru.Nama_Guru, &objGuru.Jenis_Kelamin, &objGuru.Tanggal_Lahir, &objGuru.No_Telp, &objGuru.Kelas,
	)

	if errGuru != nil {
		return res, errGuru
	}

	sqlStatementSiswa := "SELECT * FROM siswa WHERE kelas =?"

	rows, err := con.Query(sqlStatementSiswa, kelas)
	defer rows.Close()

	if err != nil {
		return res, nil
	}

	for rows.Next() {
		err = rows.Scan(&objSiswa.Id_Siswa, &objSiswa.Nama_Siswa, &objSiswa.Nis, &objSiswa.Jenis_Kelamin, &objSiswa.No_Telp, &objSiswa.Kelas)

		arrobjSiswa = append(arrobjSiswa, objSiswa)
	}

	responseKelas := Kelas{
		Id_Kelas: objKelas.Id_Kelas,
		Kelas:    objKelas.Kelas,
	}

	responseGuru := Guru{
		Id_Guru:       objGuru.Id_Guru,
		Nama_Guru:     objGuru.Nama_Guru,
		Jenis_Kelamin: objGuru.Jenis_Kelamin,
		Tanggal_Lahir: objGuru.Tanggal_Lahir,
		No_Telp:       objGuru.No_Telp,
		Kelas:         objGuru.Kelas,
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]interface{}{
		"class": responseKelas,
		"guru":  responseGuru,
		"siswa": arrobjSiswa,
	}

	return res, nil
}
