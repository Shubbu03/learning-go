package main

import "fmt"

func intSeq() func(x int) int {
	i := 0
	return func(x int) int {
		i++
		return i + x
	}
}

func activateGiftCard() func(x int) int {
	total := 1000

	useGiftCard := func(x int) int {
		total -= x

		return total
	}

	return useGiftCard
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt(9))
	fmt.Println(nextInt(8))
	fmt.Println(nextInt(7))

	newInts := intSeq()
	fmt.Println(newInts(6))

	giftCard := activateGiftCard()

	fmt.Println(giftCard(200))

	giftCard1 := activateGiftCard()
	fmt.Println(giftCard1(100))
}
