package main

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	unsortedList := []int{2, 5, 7, 3, 10, 1}
	sortedList := []int{1, 2, 3, 5, 7, 10}

	resList := selectionSort(unsortedList)

	if !reflect.DeepEqual(sortedList, resList) {
		t.Errorf("Lists are not the same. Expected %v, got %v", sortedList, resList)
	}
}
