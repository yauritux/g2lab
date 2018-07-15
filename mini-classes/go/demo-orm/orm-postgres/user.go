package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	params := mux.Vars(r)
	name := params["name"]
	email := params["email"]
	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "New user successfully created.")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	params := mux.Vars(r)
	name := params["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully deleted user with name %s", name)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5432, user=postgres dbname=test sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	params := mux.Vars(r)
	name := params["name"]
	email := params["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "Successfully updated user email to %s", email)
}
