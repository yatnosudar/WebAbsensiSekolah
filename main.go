package main

import (
	"echo/WebAbsensiSekolah/db"
	"echo/WebAbsensiSekolah/routes"
)

func main() {
	db.Init()

	e := routes.Routes()

	e.Logger.Fatal(e.Start(":2000"))
}
