package main

import (
	"sort"
)

// Given an input integer array nums, write a function to find all unique triplets [nums[i],
// nums[j], nums[k]] such that i, j, and k are distinct indices, and the sum of nums[i], nums[j],
// and nums[k] equals zero. Ensure that the resulting list does not contain any duplicate triplets.

// examples:
// input: nums = [-1,0,1,2,-1,-1]
// output: [[-1,-1,2],[-1,0,1]]

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for k := range len(nums) - 2 {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		left, right := k+1, len(nums)-1

		for left < right {
			sum := nums[k] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[k], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return result
}
