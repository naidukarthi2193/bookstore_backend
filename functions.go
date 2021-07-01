package main

import (
	"fmt"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "BookStore API using Go and Cassandra")   
}


// func CountAllStudents(w http.ResponseWriter, r *http.Request) {

// 	var Count string
// 	err := Session.Query("SELECT count(*) FROM students").Scan(&Count)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Fprintf(w, "%s ", Count)

// }

// func UpdateStudent(w http.ResponseWriter, r *http.Request) {
// 	StudentID := mux.Vars(r)["id"]
// 	var UpdateStudent Student
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data properly")
// 	}
// 	json.Unmarshal(reqBody, &UpdateStudent)
// 	if err := Session.Query("UPDATE students SET firstname = ?, lastname = ?, age = ? WHERE id = ?",
// 		UpdateStudent.Firstname, UpdateStudent.Lastname, UpdateStudent.Age, StudentID).Exec(); err != nil {
// 		fmt.Println("Error while updating")
// 		fmt.Println(err)
// 	}
// 	fmt.Fprintf(w, "updated successfully")

// }
