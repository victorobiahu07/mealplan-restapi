package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "mealplan-restapi/config"
	. "mealplan-restapi/dao"
	. "mealplan-restapi/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var config = Config{}
var dao = ControllerDAO{}

// GET list of all meal plans
func getAllMealPlans(w http.ResponseWriter, r *http.Request) {
	meals, err := dao.FindAllMeals()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, meals)
}

// GET list of customers
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	users, err := dao.FindAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

// GET a meal plan by its ID
func getMealPlanById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	meal, err := dao.FindMealById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Meal ID")
		return
	}
	respondWithJson(w, http.StatusOK, meal)
}

// GET meals plans by market
func getMealsByMarket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("market")
	fmt.Println(params["market"])
	meals, err := dao.FindAllMealsByMarket(params["market"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Meal request for market")
		return
	}
	respondWithJson(w, http.StatusOK, meals)
}

// GET meal plans by semester
func getMealBySemester(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Println(params["semester"])
	meals, err := dao.FindAllMealsByMarket(params["semester"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Meal plan in this semester")
		return
	}
	respondWithJson(w, http.StatusOK, meals)
}

// GET a customer by id ID
func getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := dao.FindUserById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJson(w, http.StatusOK, user)
}

// POST a new meal plan
func createMealPlan(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var meal Meal

	if err := json.NewDecoder(r.Body).Decode(&meal); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	meal.ID = bson.NewObjectId()
	if err := dao.Insert_meal(meal); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, meal)
}

//creates a new customer using a POST method
func createCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := dao.Insert_user(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}

// PUT function to update an existing meal
func updateMealPlan(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var meal Meal

	if err := json.NewDecoder(r.Body).Decode(&meal); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateMeal(meal); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// PUT function to update an existing meal
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Println("working")
	if err := dao.UpdateUser(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing meal plan
func deleteMealPlan(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var meal Meal
	if err := json.NewDecoder(r.Body).Decode(&meal); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.DeleteMeal(meal); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

//deletes an existing customer
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.DeleteUser(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	fmt.Println("Starting the Good Uncle Meal Plan application...")
	r.HandleFunc("/meals", getAllMealPlans).Methods("GET")
	r.HandleFunc("/users", getAllCustomers).Methods("GET")
	r.HandleFunc("/meals", createMealPlan).Methods("POST")
	r.HandleFunc("/users", createCustomer).Methods("POST")
	r.HandleFunc("/meals", updateMealPlan).Methods("PUT")
	r.HandleFunc("/users", updateCustomer).Methods("PUT")
	r.HandleFunc("/meals", deleteMealPlan).Methods("DELETE")
	r.HandleFunc("/users", deleteCustomer).Methods("DELETE")
	r.HandleFunc("/meals/{id}", getMealPlanById).Methods("GET")
	r.HandleFunc("/mealsbymarket/{market}", getMealsByMarket).Methods("GET")
	r.HandleFunc("/users/{id}", getCustomer).Methods("GET")
	http.ListenAndServe(":8080", r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
