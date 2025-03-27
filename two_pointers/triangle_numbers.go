package main

import (
	"sort"
)

// Write a function to count the number of triplets in an integer array nums that could form the
// sides of a triangle. The triplets do not need to be unique.

// examples:
// input: nums = [11,4,9,6,15,18]
// output: 10

func triangleNumbers(nums []int) int {
	count := 0
	sort.Ints(nums)

	for k := 2; k <= len(nums)-1; k++ {
		left, right := 0, k-1
		for left < right {
			if nums[left]+nums[right] > nums[k] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}
