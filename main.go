package main

import "log"

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Pocus"

	log.Println(myMap["dog"])
}