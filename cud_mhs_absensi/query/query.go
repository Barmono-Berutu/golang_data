package query

import (
	"cud_mhs_absensi/database"
)

type Users struct {
	Id    int
	Name  string
	Nim   int
	Class string
}

func GetAll() []Users {
	rows, err := database.DB.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}

	var users []Users

	for rows.Next() {
		var user Users
		err := rows.Scan(&user.Id, &user.Name, &user.Nim, &user.Class)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users
}

func AddData(user Users) bool {
	res, err := database.DB.Exec(`INSERT INTO user(name,nim,class) VALUE (?,?,?)`, user.Name, user.Nim, user.Class)

	if err != nil {
		panic(err)
	}

	li, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	return li > 0
}

func Delete(id int) error {
	_, err := database.DB.Exec(`DELETE FROM user WHERE id=?`, id)
	return err
}
