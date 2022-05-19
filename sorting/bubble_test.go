package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{2, 4, 3, 5, 8, 7}
	resList := []int{2, 3, 4, 5, 7, 8}

	sortedList := bubbleSort(list)

	if !reflect.DeepEqual(sortedList, resList) {
		t.Errorf("Lists are incorrect. Expected %v, got %v", resList, sortedList)
	}
}
