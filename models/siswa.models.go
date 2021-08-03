package models

import (
	"echo/WebAbsensiSekolah/db"
	"net/http"
)

type SiswaDetail struct {
	Id_Siswa      int    `json:"id_siswa"`
	Nama_Siswa    string `json:"nama_siswa"`
	Nis           int    `json:"nis"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	No_Telp       string `json:"no_telp"`
	Kelas         string `json:"kelas"`
}

func StoreSiswa(Nama_Siswa string, Nis int, Jenis_Kelamin string, No_telp string, Kelas string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT siswa (nama_siswa, nis, jenis_kelamin, no_telp, kelas) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama_Siswa, Nis, Jenis_Kelamin, No_telp, Kelas)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateSiswa(Nama_Siswa string, Nis int, jenis_kelamin string, No_Telp string, kelas string, Id_Siswa int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE siswa SET nama_siswa = ?, nis = ?, jenis_kelamin = ?, No_Telp = ?, kelas = ? WHERE id_siswa = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama_Siswa, Nis, jenis_kelamin, No_Telp, kelas, Id_Siswa)
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

func DeleteSiswa(Id_Siswa int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM siswa WHERE id_siswa = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Id_Siswa)
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
