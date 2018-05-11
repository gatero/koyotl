package mysql

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

//: MYSQL
func MySQL() (*gorm.DB, error) {
	MYSQL_CONNECTION_STRING := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8",
		os.Getenv("MYSQL_CONTAINER"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
	)

	return gorm.Open("mysql", MYSQL_CONNECTION_STRING)
}
