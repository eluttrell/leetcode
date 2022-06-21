package main

import "math"

func maxSubArray(nums []int) int {
	sum := math.MinInt32
	num := 0
	for _, v := range nums {
		num += v
		if num > sum {
			sum = num
		}
		if num < 0 {
			num = 0
		}
	}
	return sum
}
