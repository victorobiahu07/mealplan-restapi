package models

import "gopkg.in/mgo.v2/bson"

// Represents a meal, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document

type Semester struct {
	Session   string `bson:"session" json:"session"`
	StartDate string `bson:"start_date" json:"start_date" `
	EndDate   string `bson:"end_date" json:"end_date" `
}
type Meal struct {
	ID                    bson.ObjectId `bson:"_id" json:"id"`
	Name                  string        `bson:"name" json:"name"`
	WeeklyCost            int           `bson:"weekly_cost" json:"weekly_cost"`
	Description           string        `bson:"description" json:"description"`
	SpringSemesterDates   Semester      `bson:"spring_semester" json:"spring_semester"`
	FallSemesterDates     Semester      `bson:"fall_semester" json:"fall_semester"`
	Market                string        `bson:"market" json:"market"`
	ThirdAcademicSemester Semester      `bson:"third_semester" json:"third_semester"`
}
