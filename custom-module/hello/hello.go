package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    message := greetings.Hi("shubbu")
    fmt.Println(message)
}