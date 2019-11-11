package dao

import (
	"log"

	. "mealplan-restapi/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ControllerDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION_meal = "meals"
	COLLECTION_user = "users"
)

// Establish a connection to database
func (m *ControllerDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of meals
func (m *ControllerDAO) FindAllMeals() ([]Meal, error) {
	var meals []Meal
	err := db.C(COLLECTION_meal).Find(bson.M{}).All(&meals)
	return meals, err
}


// Find list of meals ordered by market
//newly added
func (m *ControllerDAO) FindAllMealsOrderedByMarket() ([]Meal, error) {
	// *****NOV 11TH UPDATE***
	var meals []Meal
	err := db.C(COLLECTION_meal).Find(bson.M{}).Sort("market").All(&meals)
	return meals, err
}

// Find list of meals ordered by weekly cost
//newly added
func (m *ControllerDAO) FindAllMealsOrderedByWeeklyPrice() ([]Meal, error) {
	var meals []Meal
	err := db.C(COLLECTION_meal).Find(bson.M{}).Sort("weekly_cost").All(&meals)
	return meals, err
}


// Find list of users
func (m *ControllerDAO) FindAllUsers() ([]User, error) {
	var users []User
	err := db.C(COLLECTION_user).Find(bson.M{}).All(&users)
	return users, err
}

// // Find a meal by its id
func (m *ControllerDAO) FindMealById(id string) (Meal, error) {
	var meal Meal
	err := db.C(COLLECTION_meal).FindId(bson.ObjectIdHex(id)).One(&meal)
	return meal, err
}

// // Find a meal by its market
func (m *ControllerDAO) FindAllMealsByMarket(market string) ([]Meal, error) {

	var meals []Meal
	err := db.C(COLLECTION_meal).Find(bson.M{"market": market}).All(&meals)
	return meals, err
}

// // Find a user by its id
func (m *ControllerDAO) FindUserById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION_user).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a meal into database
func (m *ControllerDAO) Insert_meal(meal Meal) error {
	err := db.C(COLLECTION_meal).Insert(&meal)
	return err
}
func (m *ControllerDAO) Insert_user(user User) error {
	err := db.C(COLLECTION_user).Insert(&user)
	return err
}

// Delete an existing meal
func (m *ControllerDAO) DeleteMeal(meal Meal) error {
	err := db.C(COLLECTION_meal).Remove(&meal)
	return err
}

// Delete an existing user
func (m *ControllerDAO) DeleteUser(user User) error {
	err := db.C(COLLECTION_user).Remove(&user)
	return err
}

// Update an existing meal
func (m *ControllerDAO) UpdateMeal(meal Meal) error {
	err := db.C(COLLECTION_meal).UpdateId(meal.ID, &meal)
	return err
}

// Update an existing user
func (m *ControllerDAO) UpdateUser(user User) error {
	err := db.C(COLLECTION_user).UpdateId(user.ID, &user)
	return err
}
