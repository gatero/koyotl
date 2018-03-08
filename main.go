package main

import (
	"app/db"
	"app/router"
	"fmt"
)

func main() {
	db, _ := db.Open()
	defer db.Close()
	router := &router.Config{DB: db}
	router.Up()

	fmt.Printf("db: %s\n", db)
}
