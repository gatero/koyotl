package db

import (
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

//: MONGO DB
func Mongo() (*mgo.Session, error) {
	MONGO_CONNECTION_PARAMS := &mgo.DialInfo{
		Timeout:  60 * time.Second,
		Addrs:    []string{os.Getenv("MONGO_CONTAINER")},
		Database: os.Getenv("MONGO_DATABASE"),
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}

	return mgo.DialWithInfo(MONGO_CONNECTION_PARAMS)
}
