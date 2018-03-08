package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MYSQL_HOST string = os.Getenv("MYSQL_CONTAINER")
var MYSQL_DATABASE string = os.Getenv("MYSQL_DATABASE")
var MYSQL_PORT string = os.Getenv("MYSQL_PORT")
var MYSQL_USERNAME string = os.Getenv("MYSQL_USERNAME")
var MYSQL_PASSWORD string = os.Getenv("MYSQL_ROOT_PASSWORD")
var MYSQL_CONNECTION_STRING string = fmt.Sprintf(
	"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8",
	MYSQL_USERNAME,
	MYSQL_PASSWORD,
	MYSQL_HOST,
	MYSQL_PORT,
	MYSQL_DATABASE,
)

func Open() (*gorm.DB, error) {
	return gorm.Open("mysql", MYSQL_CONNECTION_STRING)
}
