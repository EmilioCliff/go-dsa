package main

func variableLengthSlidingWindow(nums []int) int {
	// state := map[string]int{}
	start := 0
	maxCount := 0

	for end := range len(nums) {
		// extend window
		// add nums[end] to state in O(1) in time

		// for state not valid {
		// repeatedly contract window until it is valid again
		// remove nums[start] from state in O(1) in time
		start += 1
		// }

		// INVARIANT: state of current window is valid here.
		maxCount = max(maxCount, end-start+1)
	}

	return maxCount
}

func fixedLenghtSlidingWindow(nums []int, k int) int {
	// state := // choose appropriate data structure
	start := 0
	maxCount := 0

	for end := range len(nums) {
		// extend window
		// add nums[end] to state in O(1) in time

		if end-start+1 == k {
			// INVARIANT: size of the window is k here.
			// maxCount = max(maxCount, contents of state)

			// contract window
			// remove nums[start] from state in O(1) in time
			start += 1
		}
	}

	return maxCount
}
