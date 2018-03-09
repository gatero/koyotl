package main

import (
	"app/db"
	"app/router"
)

func main() {
	DB, _ := db.OpenMySQL()
	defer DB.Close()

	//: Router init
	router := &router.Config{DB: DB}
	router.Up()
}
