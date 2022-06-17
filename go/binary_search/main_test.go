package main

import "testing"

func TestBinarySearch(t *testing.T) {
	result1 := binarySearch([]int{-1, 1, 2, 3, 4, 5, 6, 7, 9}, 9)

	if result1 != 8 {
		t.Errorf("Expected %v for result1, got %v", 8, result1)
	}

	result2 := binarySearch([]int{-1, 1, 2, 3, 4, 5, 6, 7}, 9)

	if result2 != -1 {
		t.Errorf("Expected %v for result2, got %v", -1, result2)
	}

	result3 := binarySearch([]int{9}, 9)

	if result3 != 0 {
		t.Errorf("Expected %v for result3, got %v", 0, result3)
	}

	result4 := binarySearch([]int{-1, 9}, 9)

	if result4 != 1 {
		t.Errorf("Expected %v for result4, got %v", 1, result4)
	}
}
