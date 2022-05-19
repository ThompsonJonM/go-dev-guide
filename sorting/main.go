package main

import "fmt"

func main() {
	list := []int{1,5,7,10,15,22,28,30,35,40}

	res := binarySort(list, 10)
	fmt.Println(res)

	unsortedList := []int{2, 5, 7, 8, 1, 3, 6}

	resList := bubbleSort(unsortedList)
	fmt.Println(resList)

	unsortedSelectionList := []int{2, 5, 7, 4, 10}

	resList = selectionSort(unsortedSelectionList)
	fmt.Println(resList)
}
