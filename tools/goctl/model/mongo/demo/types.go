package model

//go:generate goctl model mongo -t User
import "github.com/globalsign/mgo/bson"

type User struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}
