package main

import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"github.com/gocql/gocql"

	"github.com/gorilla/mux"
)

// Create Purchases
func CreatePurchases(w http.ResponseWriter, r *http.Request) {
	var NewPurchases Purchases
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "wrong data")
	}
	json.Unmarshal(reqBody, &NewPurchases)
	newUuid, _ := gocql.RandomUUID()
	if err := Session.Query("INSERT INTO purchases(uuid, useruuid, bookuuid, timestamp) VALUES(?, ?, ?, ?)",
		newUuid, NewPurchases.UserUUID, NewPurchases.BookUUID, time.Now()).Exec(); err != nil {
		fmt.Println("Error while inserting Purchases")
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(NewPurchases, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

// Delete Purchases by UUID
func DeletePurchases(w http.ResponseWriter, r *http.Request) {
	PurchasesID := mux.Vars(r)["id"]
	if err := Session.Query("DELETE FROM purchases WHERE uuid = ?", PurchasesID).Exec(); err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted successfully the Purchase num %s ", PurchasesID)
}

// Delete All Purchases
func DeleteAllPurchases(w http.ResponseWriter, r *http.Request) {
	if err := Session.Query("TRUNCATE purchases").Exec(); err != nil {
		fmt.Println("Error while deleting all students")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted all successfully")
}

// Get Purchases
func GetPurchases(w http.ResponseWriter, r *http.Request) {
	var purchases []Purchases
	m := map[string]interface{}{}
	iter := Session.Query("SELECT * FROM purchases").Iter()
	for iter.MapScan(m) {
		purchases = append(purchases, Purchases{
			UUID:				m["uuid"].(gocql.UUID),
			UserUUID:			m["useruuid"].(gocql.UUID),
			BookUUID:			m["bookuuid"].(gocql.UUID),
			Timestamp:			m["timestamp"].(time.Time),
		})
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(purchases, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}
