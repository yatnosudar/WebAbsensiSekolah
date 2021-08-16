package models

import (
	"database/sql"
	"echo/WebAbsensiSekolah/db"
	"echo/WebAbsensiSekolah/helper"
	"fmt"
)

type Account struct {
	Id_Guru  int    `json:"id_guru"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func LoginGuru(username, password string) (bool, error) {
	var obj Account
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT id_guru, username, password FROM guru WHERE username = ? AND role ='guru'"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id_Guru, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match!")
		return false, err
	}

	return true, nil
}

func LoginAdmin(username, password string) (bool, error) {
	var obj Account
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT id_guru, username, password FROM guru WHERE username = ? AND role ='admin'"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id_Guru, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match!")
		return false, err
	}

	return true, nil
}

func LoginSiswa(username, password string) (bool, error) {
	var objAcc SiswaAcc

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM account_siswa"

	err := con.QueryRow(sqlStatement).Scan(
		&objAcc.Id, &objAcc.Nis, &objAcc.Password,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Please check your nis or password")
		return false, err
	}
	if err != nil {
		fmt.Println("Query error")
		return false, err
	}
	return true, nil
}
