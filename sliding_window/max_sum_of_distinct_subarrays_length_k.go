package main

// Given an integer array nums and an integer k, write a function to identify the highest possible
// sum of a subarray within nums, where the subarray meets the following criteria: its length is k,
// and all of its elements are unique.

// example:
// input: nums = [3, 2, 2, 3, 4, 6, 7, 7, -1]
// 	   k = 4
// output: 20
// Explanation: The subarrays of nums with length 4 are:
// 	   [3, 2, 2, 3] # elements 3 and 2 are repeated.
// 	   [2, 2, 3, 4] # element 2 is repeated.
// 	   [2, 3, 4, 6] # meets the requirements and has a sum of 15.
// 	   [3, 4, 6, 7] # meets the requirements and has a sum of 20.
// 	   [4, 6, 7, 7] # element 7 is repeated.
// 	   [6, 7, 7, -1] # element 7 is repeated.

func maxSumOfDistinctSubarraysLenghtK(nums []int, k int) int64 {
	seen := map[int]int{}
	maxSum, start, state := 0, 0, 0

	for end := range len(nums) {
		state += nums[end]
		seen[nums[end]] += 1

		for seen[nums[end]] > 1 {
			state -= nums[start]
			seen[nums[start]] -= 1
			start++
		}

		if end-start+1 == k {
			maxSum = max(maxSum, state)
			state -= nums[start]
			delete(seen, nums[start])
			start++
		}
	}

	return int64(maxSum)
}
