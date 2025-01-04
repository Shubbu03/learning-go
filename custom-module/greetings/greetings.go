package greetings

import (
	"fmt"
)

func Hi(name string) string {
	message := fmt.Sprintf("Hi there %v", name)

	return message
}
