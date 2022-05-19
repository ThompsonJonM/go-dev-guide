package main

func selectionSort(list []int) []int {
	for i := 0; i < len(list) - 1; i++ {
		lowest := i

		for j := i + 1; j < len(list); j++ {
			if list[j] < list[lowest] {
				lowest = j
			}
		}

		if lowest != i {
			list[i], list[lowest] = list[lowest], list[i]
		}
	}

	return list
}
