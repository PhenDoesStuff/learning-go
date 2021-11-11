package main

import "log"

var s = "seven"

func main() {
	var s2 = "six"
	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx") // Which s is this going to refer to?
}

func saySomething(s string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s, "World"
}