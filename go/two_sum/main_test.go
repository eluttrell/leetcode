package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 14, 19, 10, 7}
  expected := []int{1, 7}
	twoSum := TwoSum(arr, 9)
  noMatch := false
  for i, v := range twoSum {
    if i > len(expected) || v != expected[i] {
      noMatch = true
    }
  }
	if noMatch {
		t.Errorf("result should be %v, got %v", expected, twoSum)
	}
}
