package greetings

import "fmt"

// Baş harfi büyük olması public olmasını sağlıyor
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %s. Welcome!", name)
	return message
}
