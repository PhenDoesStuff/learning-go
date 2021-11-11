package main

import "log"

var s = "seven"

func main() {
	var s2 = "six"

	s := "eight"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx") // Which s is this going to refer to?
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "World"
}