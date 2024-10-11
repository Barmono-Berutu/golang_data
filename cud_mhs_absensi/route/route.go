package route

import (
	"cud_mhs_absensi/query"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./public/index.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	data := query.GetAll()
	err = temp.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
func Add(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		temp, err := template.ParseFiles("./public/add.html")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user query.Users

		user.Name = r.FormValue("name")
		// Konversi string nim ke integer
		if nim, err := strconv.Atoi(r.FormValue("nim")); err == nil {
			user.Nim = nim
		} else {
			http.Error(w, "Invalid NIM value", http.StatusBadRequest)
			return
		}
		user.Class = r.FormValue("class")

		ok := query.AddData(user)

		if !ok {
			temp, _ := template.ParseFiles("./public/add.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	err = query.Delete(id)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
