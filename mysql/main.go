package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func main() {
	// Open DB
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/sea_battle")
	if err != nil {
		panic(err)
	}
	defer func() { err = db.Close() }()

	// Insert data
	rows, err := db.Query("INSERT INTO `users` (`name`,`age`) VALUES ('Alex', 35), ('Bob', 27)")
	if err != nil {
		panic(err)
	}
	defer func() { err = rows.Close() }()

	// Select data
	res, err := db.Query("SELECT `name`,`age` FROM `users`")
	if err != nil {
		panic(err)
	}
	defer func() { err = res.Close() }()

	for res.Next() {
		var user User
		if err = res.Scan(&user.Name, &user.Age); err != nil {
			panic(err)
		}
		_, _ = fmt.Fprintf(os.Stdout, "Use: %s, age: %d\n", user.Name, user.Age)
	}
}
