package main

import (
	"log"

	"github.com/stephenmontague/learning-go/helpers"
)

const numPool = 10000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)
	
	num := <- intChan

	log.Println(num)
}