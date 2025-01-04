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

func Hellos(names []string) (map[string]string,error){
	messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hi(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}
