package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	MealID bson.ObjectId `bson:"meal_id" json:"meal_id"`
}
