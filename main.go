package main

import (
	"app/db"
	"app/router"
)

func main() {
	mysql := &db.Mysql{}
	DB, _ := mysql.Open()
	defer DB.Close()

	//: Router init
	router := &router.Config{DB: DB}
	router.Up()
}
