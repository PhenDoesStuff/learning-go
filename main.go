package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "John@Smith.com", 30})
	users = append(users, User{"Mary", "Jones", "Mary@Jones.com", 25})
	users = append(users, User{"Sally", "Brown", "Sally@Brown.com", 12})
	users = append(users, User{"Alex", "Anderson", "Alex@Anderson.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}