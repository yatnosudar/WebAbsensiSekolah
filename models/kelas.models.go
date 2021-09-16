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
	Id_Guru  int    `json:"id_guru"`
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
}

func GetListKelas() (Response, error) {
	var obj Kelas
	var arrobj []Kelas
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM kelas ORDER BY id_kelas"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Kelas, &obj.Kelas, &obj.Id_Guru)
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

	// perintah sql untuk memilih 1 kelas
	sqlStatementKelas := "SELECT * FROM kelas WHERE kelas = ?"

	errKelas := con.QueryRow(sqlStatementKelas, kelas).Scan(
		&objKelas.Id_Kelas, &objKelas.Kelas, &objKelas.Id_Guru,
	)

	if errKelas == sql.ErrNoRows {
		fmt.Println("Kelas Tidak Ditemukan")
		return res, errKelas
	}

	if errKelas != nil {
		return res, errKelas
	}

	// menampilkan data guru yang menjadi wali kelas
	sqlStatementGuru := "SELECT id_guru, nama_guru, jenis_kelamin, tanggal_lahir, no_telp FROM guru WHERE id_guru =?"

	errGuru := con.QueryRow(sqlStatementGuru, &objKelas.Id_Guru).Scan(
		&objGuru.Id_Guru, &objGuru.Nama_Guru, &objGuru.Jenis_Kelamin, &objGuru.Tanggal_Lahir, &objGuru.No_Telp,
	)

	if errGuru != nil {
		return res, errGuru
	}

	// menampilkan data siswa yang ada di kelas tersebut
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
		Id_Guru:  objKelas.Id_Guru,
	}

	responseGuru := Guru{
		Id_Guru:       objGuru.Id_Guru,
		Nama_Guru:     objGuru.Nama_Guru,
		Jenis_Kelamin: objGuru.Jenis_Kelamin,
		Tanggal_Lahir: objGuru.Tanggal_Lahir,
		No_Telp:       objGuru.No_Telp,
	}

	// response
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]interface{}{
		"kelas": responseKelas,
		"guru":  responseGuru,
		"siswa": arrobjSiswa,
	}

	return res, nil
}

func AddKelas(kelas string, id_guru int) (Response, error) {
	var res Response

	con := db.CreateCon()

	// perintah sql untuk menambahkan kelas
	sqlStatement := "INSERT kelas (kelas, id_guru) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(kelas, id_guru)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	fmt.Println("Last insert id : ", lastInsertedId)

	// response
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]string{
		"pesan": "anda telah berhasil menambahkan kelas",
	}

	return res, nil
}

func UpdateKelas(kelas string, id_guru int, id_kelas int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE kelas SET kelas = ?, id_guru = ? WHERE id_kelas = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(kelas, id_guru, id_kelas)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteKelas(Id_Kelas int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM kelas WHERE id_kelas = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Id_Kelas)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
