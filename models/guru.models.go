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

func FetchAllGuru() (Response, error) {
	var objGuru GuruDetail
	var arrobjGuru []GuruDetail
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM guru"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&objGuru.Id_Guru, &objGuru.Nama_Guru, &objGuru.Jenis_Kelamin, &objGuru.Tanggal_Lahir, &objGuru.No_Telp)
		arrobjGuru = append(arrobjGuru, objGuru)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobjGuru

	return res, nil
}

func StoreGuru(Nama_Guru string, Jenis_Kelamin string, Tanggal_Lahir string, No_telp string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT guru (nama_guru, jenis_kelamin, tanggal_lahir, no_telp) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_telp)
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

func UpdateGuru(Nama_Guru string, Jenis_Kelamin string, Tanggal_Lahir string, No_Telp string, Id_Guru int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE guru SET nama_guru = ?, jenis_kelamin = ?, tanggal_lahir = ?, No_Telp = ? WHERE id_guru = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama_Guru, Jenis_Kelamin, Tanggal_Lahir, No_Telp, Id_Guru)
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
