package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	router.HandleFunc("/books", CreateBook).Methods("POST")              
	router.HandleFunc("/books", GetBook).Methods("GET")
	// router.HandleFunc("/books/{id}", UpdateBook).Methods("PATCH") 
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	router.HandleFunc("/books", DeleteAllBook).Methods("DELETE")  

	router.HandleFunc("/purchases", CreatePurchases).Methods("POST")              
	router.HandleFunc("/purchases/{id}", GetPurchases).Methods("GET")
	router.HandleFunc("/purchases/{id}", DeletePurchases).Methods("DELETE")     
	router.HandleFunc("/purchases", DeleteAllPurchases).Methods("DELETE")   

	router.HandleFunc("/users", CreateUser).Methods("POST")              
	router.HandleFunc("/users", CheckUser).Methods("GET")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")   
	router.HandleFunc("/users", DeleteAllUsers).Methods("DELETE")     


	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))

}
