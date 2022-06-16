package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	max := 0
	lowest := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] < lowest {
			lowest = prices[i]
		} else if prices[i]-lowest > max {
			max = prices[i] - lowest
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
