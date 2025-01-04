package main

import "fmt"

func main() {
	// n := 10
	// i := 1
	// for i := 0; i < n; i++ {
	// 	fmt.Printf("GO in loop for %v \n", i)
	// }
	// sum := 0
	// for i < 3 {
	// 	fmt.Println("Sum is", sum+i)
	// 	i++
	// }

	// for j := range 3 {
	// 	fmt.Println("range", j)
	// }

	// for {
	// 	fmt.Println("loop")
	// 	break
	// }

	// for n := range 6 {
	// 	if n%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	var a, b, c int = 0, 0, 0

	fmt.Print("Enter three numbers::\n")
	fmt.Scan(&a, &b, &c)

	if a > b && a > c {
		fmt.Printf("Greatest number is %v", a)
	} else if b > a && b > c {
		fmt.Printf("Greatest number is %v", b)
	} else {
		fmt.Printf("Greatest number is %v", c)
	}

	fmt.Print("Execution complete!!")
}
