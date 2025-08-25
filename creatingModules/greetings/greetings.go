package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {

	// error handling
	if name == "" {
		return "", errors.New("empty name")
	}

    // message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat())	
    return message, nil // nil - no error
}


func Hellos(names []string) (map[string]string, error) {
	// return greetings for multiple people
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}


func randomFormat() string {
	
	// return a random greeting from one of the greeting formats
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you %v",
		"What is up %v!",
	}

	return formats[rand.Intn(len(formats))]
}