package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	user := User {
		FirstName: "Stephen",
		LastName: "Montague",
		PhoneNumber: "123-456-7890",
		Age: 26,
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}