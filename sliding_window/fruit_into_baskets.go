package main

// Write a function to calculate the maximum number of fruits you can collect from an integer array
// fruits, where each element represents a type of fruit. You can start collecting fruits from any
// position in the array, but you must stop once you encounter a third distinct type of fruit. The
// goal is to find the longest subarray where at most two different types of fruits are collected.

// example:
// input: fruits = [3, 3, 2, 1, 2, 1, 0]
// output: 4
// explanation: We can pick up 4 fruit from the subarray [2, 1, 2, 1]

func fruitIntoBaskets(fruits []int) int {
	state := map[int]int{}
	start := 0
	maxCount := 0

	for end := range len(fruits) {
		state[fruits[end]] = state[fruits[end]] + 1

		for len(state) > 2 {
			state[fruits[start]] = state[fruits[start]] - 1
			if state[fruits[start]] == 0 {
				delete(state, fruits[start])
			}
			start++
		}

		maxCount = max(maxCount, end-start+1)
	}

	return maxCount
}
