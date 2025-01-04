package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//WITHOUT ERROR HANDLING =>
// func Hi(name string) string {
// 	message := fmt.Sprintf("Hi there %v", name)

// 	return message
// }

// WITH ERROR HANDLING =>
// func Hi(name string) (string, error) {
// 	if name == "" {
// 		return "", errors.New("empty name")
// 	}
// 	message := fmt.Sprintf("Hi there %v", name)

// 	return message, nil
// }

//Random data =>
func Hi(name string) (string, error) {
    if name == "" {
        return name, errors.New("empty name")
    }
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return formats[rand.Intn(len(formats))]
}
