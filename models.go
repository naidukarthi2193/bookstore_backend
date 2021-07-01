package main

import (
	"time"
	"github.com/gocql/gocql"
)

type Book struct {
	UUID gocql.UUID  `json:"uuid"`
	Thumbnail string `json:"thumbnail"`
	Name  string `json:"name"`
	Description string `json:"description"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	Ratings int `json:"ratings"`
	Price int `json:"price"`
	AvailableQuantity int   `json:"availableQuantity"`
}

type User struct {
	UUID gocql.UUID  `json:"uuid"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Seller bool `json:"seller"`
}

type Purchases struct {
	UUID gocql.UUID  `json:"uuid"`
	UserUUID gocql.UUID  `json:"useruuid"`
	BookUUID gocql.UUID  `json:"bookuuid"`
	Timestamp  time.Time `json:"timestamp"`
}
