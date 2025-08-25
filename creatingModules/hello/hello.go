package main

import (
    "fmt"
    "example.com/greetings"
	"log"
)

func main() {

	// Set properties of the predefined Logger
    log.SetPrefix("greetings: ") // log entry prefix
    // log.SetFlags(0) // flag to disable printing the time, source file, and line number.

    message, err := greetings.Hello("Parth")
	
	// handle error
	if(err != nil) {
		log.Fatal(err)
	}

    fmt.Println(message)
}

// greetings: 2025/08/25 12:45:06 empty name
// exit status 1