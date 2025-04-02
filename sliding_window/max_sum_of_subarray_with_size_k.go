package main

// Given an array of integers nums and an integer k, find the maximum sum of any contiguous subarray
// of size k.

// example:
// input: nums = [2, 1, 5, 1, 3, 2], k = 3
// output: 9
// Explanation: The subarray with the maximum sum is [5, 1, 3] with a sum of 9.

func maxSumOfSubarrayWithSizeK(nums []int, k int) int {
	state, start, maxSum := 0, 0, 0

	for end := range len(nums) {
		state += nums[end]

		if end-start+1 == k {
			maxSum = max(maxSum, state)
			state -= nums[start]
			start++
		}
	}

	return maxSum
}
