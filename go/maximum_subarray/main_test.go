package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if result != 6 {
		t.Errorf("Expected %d but got %d", 6, result)
	}
}
