package models

import (
	"echo/WebAbsensiSekolah/db"
	"net/http"
)

type SiswaDetail struct {
	Nama_Siswa    string `json:"nama_siswa"`
	Nis           int    `json:"nis"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	No_Telp       string `json:"no_telp"`
	Kelas         string `json:"kelas"`
}

func FetchAllSiswa() (Response, error) {
	var obj SiswaDetail
	var arrobj []SiswaDetail
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM siswa"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Nama_Siswa, &obj.Nis, &obj.Jenis_Kelamin, &obj.No_Telp, &obj.Kelas)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res, nil
}

func StoreSiswa(Nis int, Nama_Siswa string, Jenis_Kelamin string, No_telp string, Kelas string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT siswa (nis, nama_siswa, jenis_kelamin, no_telp, kelas) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nis, Nama_Siswa, Jenis_Kelamin, No_telp, Kelas)
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

func UpdateSiswa(Nis int, jenis_kelamin string, No_Telp string, kelas string, Nama_Siswa string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE siswa SET nis = ?, jenis_kelamin = ?, No_Telp = ?, kelas = ? WHERE nama_siswa = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nis, jenis_kelamin, No_Telp, kelas, Nama_Siswa)
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

func DeleteSiswa(Nama_Siswa string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM siswa WHERE nama_siswa = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nama_Siswa)
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
