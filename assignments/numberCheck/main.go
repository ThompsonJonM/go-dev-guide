package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenOrOdd(list)
}

func evenOrOdd(numbers []int) {
	for _, num := range numbers {
		switch {
		case num % 2 == 0:
			fmt.Printf("%d is even\n", num)
		default:
			fmt.Printf("%d is odd\n", num)
		}
	}
}
