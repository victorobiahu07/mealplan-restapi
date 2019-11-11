package models

import "gopkg.in/mgo.v2/bson"


//every user has an id and meal plan ID, I included the name property as a bonus
type User struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	MealID bson.ObjectId `bson:"meal_id" json:"meal_id"`
}
