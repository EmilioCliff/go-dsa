package main

// Write a function to sort a given integer array nums in-place (and without the built-in sort
// function), where the array contains n integers that are either 0, 1, and 2 and represent the
// colors red, white, and blue. Arrange the objects so that same-colored ones are adjacent, in the
// order of red, white, and blue (0, 1, 2).

// examples:
// input: [2,1,2,0,1,0,1,0,1]
// output: [0,0,0,1,1,1,1,2,2]

func sortColors(nums []int) []int {
	for k := range len(nums) {
		if k == 0 {
			continue
		}

		left := 0
		for left < k {
			if nums[k] < nums[left] {
				nums[left], nums[k] = nums[k], nums[left]
			}
			left++
		}
	}

	return nums
}
