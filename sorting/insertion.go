package main

func insertionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := i; j > 0 && list[j-1] > list[j]; j-- {
			list[j], list[j-1] = list[j-1], list[j]
		}
	}

	return list
}
