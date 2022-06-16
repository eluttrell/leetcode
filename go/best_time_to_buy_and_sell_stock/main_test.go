package main

import "testing"

func TestMaxProfit(t *testing.T) {
	maxProfit := maxProfit([]int{7, 1, 5, 3, 6, 4})

	if maxProfit != 5 {
		t.Errorf("Error calculating max profit, expected %v, got %v", 5, maxProfit)
	}
}
