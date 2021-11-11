package main

import "log"

func main() {
	var mySlice []int

	mySlice = append(mySlice, 5)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	log.Println(mySlice)
}