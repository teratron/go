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
	var err error

	// Open DB
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sea_battle")
	if err != nil {
		panic(err)
	}
	defer func() { err = db.Close() }()
	fmt.Println(db)

	// Insert data
	rows, err := db.Query("INSERT INTO `users` (`name`,`age`) VALUES ('Alex', 25)")
	if err != nil {
		panic(err)
	}
	defer func() { err = rows.Close() }()
	fmt.Println(rows)

	// Select data
	res, err := db.Query("SELECT `name`,`age` FROM `users`")
	if err != nil {
		panic(err)
	}
	defer func() { err = res.Close() }()

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		_, _ = fmt.Fprintf(os.Stdout, "Use: %s, age: %d", user.Name, user.Age)
	}
}
