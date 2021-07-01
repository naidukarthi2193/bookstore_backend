package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

// Create User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var NewUser User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "wrong data")
	}
	fmt.Println(string(reqBody))
	err1 := json.Unmarshal([]byte(reqBody), &NewUser)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	fmt.Println(NewUser)
	NewUser.UUID, _ = gocql.RandomUUID()
	if err := Session.Query("INSERT INTO users (uuid, name, email, password, seller) VALUES (?, ?, ?, ?, ?)",
		NewUser.UUID, NewUser.Name, NewUser.Email, NewUser.Password, NewUser.Seller).Exec(); err != nil {
		fmt.Println("Error while inserting User")
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(NewUser, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

// Delete Users by UUID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	UserID := mux.Vars(r)["id"]
	if err := Session.Query("DELETE FROM users WHERE id = ?", UserID).Exec(); err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted successfully the User num %s ", UserID)
}

// Delete All Users
func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	if err := Session.Query("TRUNCATE users").Exec(); err != nil {
		fmt.Println("Error while deleting all students")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted all successfully")
}

// Get Users
func CheckUser(w http.ResponseWriter, r *http.Request) {
	var users []User
	m := map[string]interface{}{}
	iter := Session.Query("SELECT * FROM users").Iter()
	for iter.MapScan(m) {
		users = append(users, User{
			UUID:			m["uuid"].(gocql.UUID),
			Name:			m["name"].(string),
			Email:			m["email"].(string),
			Password:		m["password"].(string),
			Seller:			m["seller"].(bool),
		})
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(users, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}
