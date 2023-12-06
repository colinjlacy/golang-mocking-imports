package main

import (
	"github.com/colinjlacy/golang-mocking-imports/create"
	"log"
)

func main() {
	if err := create.CreateFile("this.txt", "that, those, them"); err != nil {
		log.Fatal(err)
	}
}
