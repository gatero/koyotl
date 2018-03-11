package db

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	mgo "gopkg.in/mgo.v2"
)

//: MYSQL

type Mysql struct {
	Host     string
	Database string
	Port     string
	Username string
	Password string
}

func (m *Mysql) Open() (*gorm.DB, error) {
	m.Host = os.Getenv("MYSQL_CONTAINER")
	m.Database = os.Getenv("MYSQL_DATABASE")
	m.Port = os.Getenv("MYSQL_PORT")
	m.Username = os.Getenv("MYSQL_USERNAME")
	m.Password = os.Getenv("MYSQL_PASSWORD")

	MYSQL_CONNECTION_STRING := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8",
		m.Host,
		m.Database,
		m.Port,
		m.Username,
		m.Password,
	)

	return gorm.Open("mysql", MYSQL_CONNECTION_STRING)
}

//: MONGO DB

type Mongo struct {
	Container string
}

func (m *Mongo) Open() (*mgo.Session, error) {
	m.Container = os.Getenv("MONGO_CONTAINER")

	MONGO_CONNECTION_PARAMS := &mgo.DialInfo{
		Addrs:    []string{m.Container},
		Timeout:  60 * time.Second,
		Database: os.Getenv("MONGO_DATABASE"),
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}

	return mgo.DialWithInfo(MONGO_CONNECTION_PARAMS)
}
