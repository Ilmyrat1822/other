package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string
	Age  uint16
}

func main() {

	db, err := sql.Open("mysql", "root:sanly2024@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users (name,age) Values('Max',25)")
	if err != nil {
		panic(err)
	}

	defer insert.Close()

	res, err := db.Query("SELECT name,age FROM users")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		println(user.Name, user.Age)
	}
}
