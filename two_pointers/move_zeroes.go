package main

// Given an integer array nums, write a function to rearrange the array by moving all zeros to the
// end while keeping the order of non-zero elements unchanged. Perform this operation in-place
// without creating a copy of the array.

// example:
// imput: [2,0,4,0,9]
// output: [2,4,9,0,0]

func moveZeroes(nums []int) []int {
	for k := range len(nums) - 1 {
		if nums[k] != 0 {
			continue
		}

		right := k + 1
		for right < len(nums) {
			if nums[right] == 0 {
				right++

				continue
			} else {
				nums[k], nums[right] = nums[right], nums[k]
				break
			}
		}
	}

	return nums
}
