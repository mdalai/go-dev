package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set Logger properties: 
	// Add prefix
	log.SetPrefix("greetings >> ")
	// Disable logging the time, src, line #
	log.SetFlags(0)


	//message, err := greetings.Hello("Gladys")
	message, err := greetings.Hello("")
	// handle error: print to console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}