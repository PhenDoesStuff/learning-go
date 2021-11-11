package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to:", i)

	whatWasSaid, otherThing := saySomething()

	fmt.Println("The function returned:", whatWasSaid, otherThing)
}

func saySomething() (string, string) {
	return "something", "something else"
}