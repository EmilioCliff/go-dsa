package main

// Given an integer input array heights representing the heights of vertical lines, write a function
// that returns the maximum area of water that can be contained by two of the lines (and the
// x-axis). The function should take in an array of integers and return an integer. Given an integer
// input array heights representing the heights of vertical lines, write a function that returns the
// maximum area of water that can be contained by two of the lines (and the x-axis). The function
// should take in an array of integers and return an integer.

// examples:
// input: heights = [3,4,1,2,2,4,1,3,2]
// output: 21
// input: heights = [1,2,1]
// output: 2

func containerWithMostWater(heights []int) int {
	maxSum := 0
	left, right := 0, len(heights)-1

	for left < right {
		w := right - left
		maxSum = max(maxSum, w*min(heights[left], heights[right]))
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxSum
}
