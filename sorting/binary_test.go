package main

import "testing"

func TestBinarySort(t *testing.T) {
	list := []int{1, 32, 5, 15, 22, 7, 16, 25, 23}

	tcs := []struct{
		Target int
		Result bool
	}{
		{
			Target: 23,
			Result: true,
		},
		{
			Target: 12,
			Result: false,
		},
	}

	for _, tc := range tcs {
		res := binarySort(list, tc.Target)
		if res != tc.Result {
			t.Errorf("expected %t but got %t", tc.Result, res)
		}
	}
}
