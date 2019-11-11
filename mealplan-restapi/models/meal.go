package models

import "gopkg.in/mgo.v2/bson"

// Represents a meal, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document

type Semester struct {
	Session   string `bson:"session" json:"session"`
	StartDate string `bson:"start_date" json:"start_date" `
	EndDate   string `bson:"end_date" json:"end_date" `
}

//included third trimester JSON property to enable inclusion of extra trimester in schools with different academic calendars
type Meal struct {
	ID                    bson.ObjectId `bson:"_id" json:"id"`
	Name                  string        `bson:"name" json:"name"`
	WeeklyCost            int           `bson:"weekly_cost" json:"weekly_cost"`
	Description           string        `bson:"description" json:"description"`
	SpringSemesterDates   Semester      `bson:"spring_semester" json:"spring_semester"`
	FallSemesterDates     Semester      `bson:"fall_semester" json:"fall_semester"`
	Market                string        `bson:"market" json:"market"`
	ThirdTrimester        Semester      `bson:"third_trimester" json:"third_trimester"`
}


/*Sample JSON for testing would look like this
{
	"id":"5dc57d4aa0401e4cd49b5188",
	"name":"Plan 1",
	"weekly_cost":50,
	"description":"Testing Meal Plan Changed Description",
	"market": "Syracuse", 
	"spring_semester_dates":{
		"session":"Spring",
		"start_date":"1-1-19",
		"end_date":"5-30-19"
	},"fall_semester_dates":{
		"session":"Fall",
		"start_date":"9-1-19",
		"end_date":"12-5-19"
	}
}*/

/*this includes trimesters
{
	"id":"1a234bc57d4",
	
	"name":"Plan 2 Gluten Free",
	"weekly_cost":60,
	"description":"Testing Meal Plan for summer trimester",
	"market": "Drexel", 
	"spring_semester_dates":{
		"session":"Spring",
		"start_date":"1-8-19",
		"end_date":"5-30-19"
	},"fall_semester_dates":{
		"session":"Fall",
		"start_date":"9-6-19",
		"end_date":"12-21-19"
	},"third_trimester_dates":{
		"session":"Summer",
		"start_date":"6-6-19",
		"end_date":"8-1-19"
	}
}
*/
