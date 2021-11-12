package main

import (
	"log"
)

func main() {
	animals := make(map[string]string)

	animals["dog"] = "fido"
	animals["cat"] = "fluffy"

	for _, animal := range animals {
		log.Println(animal)
	}
}