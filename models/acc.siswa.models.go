package models

import (
	"echo/WebAbsensiSekolah/db"
	"net/http"
)

type SiswaAcc struct {
	Id       int    `json:"id"`
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
		err = rows.Scan(&obj.Id, &obj.Nis, &obj.Password)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res, nil
}

func StoreSiswaAcc(Id int, Nis int, Password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT acc_siswa (id, nis, password) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Id, Nis, Password)
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

func UpdateSiswaAcc(Nis int, Password string, Id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE acc_siswa SET password = ?, nis = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Password, Nis, Id)
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

func DeleteSiswaAcc(Id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM acc_siswa WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(Id)
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
