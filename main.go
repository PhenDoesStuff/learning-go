package main

import "log"

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Stephen")
	mySlice = append(mySlice, "Stephen")
	mySlice = append(mySlice, "Stephen")

	log.Println(mySlice)
}