package main

import (
	"cud_mhs_absensi/database"
	"cud_mhs_absensi/route"
	"net/http"
)

func main() {
	// connecting to server
	database.ConnectDB()
	http.HandleFunc("/", route.Home)
	http.HandleFunc("/add", route.Add)
	http.HandleFunc("/delete", route.Delete)
	http.ListenAndServe(":5000", nil)
}
