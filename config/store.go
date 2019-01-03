package config

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type store struct {
	Image *mgo.Collection
}

var Store store

func InitMongo() {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		fmt.Println("error connecting to mongo", err)
	}

	db := session.DB("warehouse")
	Store.Image = db.C("image_object")
}