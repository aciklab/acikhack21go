package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Ali")
	fmt.Println(message)
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %s. Welcome!", name)
	return message
}

/*
go run main.go
go run .
go build main.go
go build .

*/
