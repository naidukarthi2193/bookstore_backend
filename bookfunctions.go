package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

// Create Books
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var NewBook Book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "wrong data")
	}
	fmt.Println(string(reqBody))
	json.Unmarshal(reqBody, &NewBook)
	fmt.Println(NewBook)
	NewBook.UUID, _ = gocql.RandomUUID()
	if err := Session.Query("INSERT INTO books (uuid, thumbnail, name, description, author, genre, ratings, price, availableQuantity) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		NewBook.UUID, NewBook.Thumbnail, NewBook.Name, NewBook.Description, NewBook.Author, NewBook.Genre, NewBook.Ratings, NewBook.Price, NewBook.AvailableQuantity).Exec(); err != nil {
		fmt.Println("Error while inserting Book")
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(NewBook, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

// Delete Books by UUID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	BookID := mux.Vars(r)["id"]
	if err := Session.Query("DELETE FROM books WHERE uuid = ?", BookID).Exec(); err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted successfully the Book num %s ", BookID)
}

// Delete All Books
func DeleteAllBook(w http.ResponseWriter, r *http.Request) {
	if err := Session.Query("TRUNCATE books").Exec(); err != nil {
		fmt.Println("Error while deleting all students")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted all successfully")

}

// Get Books
func GetBook(w http.ResponseWriter, r *http.Request) {
	var books []Book
	m := map[string]interface{}{}
	iter := Session.Query("SELECT * FROM books").Iter()
	for iter.MapScan(m) {
		fmt.Println(books)
		fmt.Println(m)
		books = append(books, Book{
			UUID:				m["uuid"].(gocql.UUID),
			Thumbnail:			m["thumbnail"].(string),
			Name:				m["name"].(string),
			Description:		m["description"].(string),
			Author:				m["author"].(string),
			Genre:				m["genre"].(string),
			Ratings:			m["ratings"].(int),
			Price:				m["price"].(int),
			AvailableQuantity:	m["availablequantity"].(int),
		})
		fmt.Println(books)
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(books, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}
