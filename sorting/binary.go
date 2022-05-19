package main

import (
	"sort"
)

func binarySort(list []int, target int) bool {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	lowerEnd := 0
	upperEnd := list[len(list) -1]

	for lowerEnd <= upperEnd {
		middlePoint := lowerEnd + upperEnd % 2
		valueAtMiddlePoint := list[middlePoint]

		switch {
		case valueAtMiddlePoint < target:
			lowerEnd = middlePoint + 1
		case valueAtMiddlePoint > target:
			upperEnd = middlePoint - 1
		case valueAtMiddlePoint == target:
			return true
		}
	}

	return false
}
