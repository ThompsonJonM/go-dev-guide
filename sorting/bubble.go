package main

func bubbleSort(list []int) []int {
	var sorted bool

	unsortedUntilIndex := len(list) - 1

	for !sorted {
		sorted = true

		for i := 0; i < unsortedUntilIndex; i++ {
			if list[i] > list[i + 1] {
				list[i], list[i + 1] = list[i + 1], list[i]

				sorted = false
			}
		}

		unsortedUntilIndex--
	}

	return list
}
