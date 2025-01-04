package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello world!!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	message:= greetings.Hi("shubbu")
	fmt.Println(message)
}
