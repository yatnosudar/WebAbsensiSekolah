package models

import (
	"echo/WebAbsensiSekolah/db"
	"net/http"
)

type SiswaAcc struct {
	Nis      int    `json:"nis"`
	Password string `json:"password"`
}

func FetchAllSiswaAcc() (Response, error) {
	var obj SiswaAcc
	var arrobj []SiswaAcc
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM acc_siswa"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Nis, &obj.Password)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res, nil
}

func StoreSiswaAcc(Nis int, Password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT acc_siswa (nis, password) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nis, Password)
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

func UpdateSiswaAcc(Nis int, Password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE acc_siswa SET password = ? WHERE nis = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Password, Nis)
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

func DeleteSiswaAcc(Nis int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM acc_siswa WHERE nis = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Nis)
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
