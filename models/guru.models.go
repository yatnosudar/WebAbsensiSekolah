package models

import (
	"echo/WebAbsensiSekolah/db"
	"net/http"
)

type GuruDetail struct {
	Id_Guru       int    `json:"id_guru"`
	Nama_Guru     string `json:"nama_guru"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Tanggal_Lahir string `json:"tanggal_lahir"`
	No_Telp       string `json:"no_telp"`
}

func StoreGuru(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_telp, username, password, role string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT guru (nama_guru, jenis_kelamin, tanggal_lahir, no_telp, username, password, role) VALUES (?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_telp, username, password, role)
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

func UpdateGuru(Nama_Guru string, Jenis_Kelamin, Tanggal_Lahir, No_Telp, username, password, role string, Id_Guru int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE guru SET nama_guru = ?, jenis_kelamin = ?, tanggal_lahir = ?, No_Telp = ?, username = ?, password = ?, role = ? WHERE id_guru = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_Telp, username, password, role, Id_Guru)
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

func DeleteGuru(Id_Guru int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM guru WHERE id_guru = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Id_Guru)
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
