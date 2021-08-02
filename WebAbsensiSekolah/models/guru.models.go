package models

import (
	"echo/WebAbsensiSekolah/db"
	"net/http"
)

type GuruDetail struct {
	Nama_Guru     string `json:"nama_guru"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Tanggal_Lahir string `json:"tanggal_lahir"`
	No_Telp       string `json:"no_telp"`
	Kelas         string `json:"kelas"`
}

func FetchAllGuru() (Response, error) {
	var obj GuruDetail
	var arrobj []GuruDetail
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM guru"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Nama_Guru, &obj.Jenis_Kelamin, &obj.Tanggal_Lahir, &obj.No_Telp, &obj.Kelas)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res, nil
}

func StoreGuru(Nama_Guru string, Jenis_Kelamin string, Tanggal_Lahir string, No_telp string, Kelas string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT guru (nama_guru, jenis_kelamin, tanggal_lahir, no_telp, kelas) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_telp, Kelas)
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

func UpdateGuru(jenis_kelamin string, Tanggal_Lahir string, No_Telp string, kelas string, Nama_Guru string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE guru SET jenis_kelamin = ?, tanggal_lahir = ?, No_Telp = ?, kelas = ? WHERE nama_guru = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(jenis_kelamin, Tanggal_Lahir, No_Telp, kelas, Nama_Guru)
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

func DeleteGuru(Nama_Guru string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM guru WHERE nama_guru = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama_Guru)
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
