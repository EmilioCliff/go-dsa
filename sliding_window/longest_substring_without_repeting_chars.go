package main

// Write a function to return the length of the longest substring in a provided string s where all
// characters in the substring are distinct.

// example:
// input: s = "eghghhgg"
// output: 3

// input: s = "substring"
// output: 8

func longestSubstringWithoutRepeatingChars(s string) int {
	state := map[string]int{}
	start := 0
	maxCount := 0

	for end := range s {
		char := string(s[end])
		state[char] += 1

		for state[char] > 1 {
			state[string(s[start])] -= 1
			start++
		}

		maxCount = max(maxCount, end-start+1)
	}

	return maxCount
}

func longestSubstringWithoutRepeatingCharsOpt(s string) int {
	state := map[string]int{}
	start := 0
	maxCount := 0

	for end := range s {
		if _, ok := state[string(s[end])]; ok {
			start = max(start, state[string(s[end])]+1)
		}

		state[string(s[end])] = end
		maxCount = max(maxCount, end-start+1)
	}

	return maxCount
}
