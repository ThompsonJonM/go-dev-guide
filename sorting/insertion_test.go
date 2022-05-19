package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	list := []int{5, 3, 6, 17, 22, 8, 4}
	sortedList := []int{3, 4, 5, 6, 8, 17, 22}

	resList := insertionSort(list)

	if !reflect.DeepEqual(sortedList, resList) {
		t.Errorf("Expected %v, got %v", sortedList, resList)
	}
}
