package main

// Write a function to find the length of the longest substring containing the same letter in a
// given string s, after performing at most k operations in which you can choose any character of
// the string and change it to any other uppercase English letter.

// example:
// input: s = "BBABCCDD"
// 	   k = 2
// output: 5
// Explanation: Replace the first 'A' and 'C' with 'B' to form "BBBBBCCDD". The longest substring
// with identical letters is "BBBBB", which has a length of 5.

func longestRepeatingCharReplacement(s string, k int) int {
	state := map[string]int{}
	maxFreq := 0
	maxLength := 0
	start := 0

	for end := range s {
		c := string(s[end])
		state[c] += 1
		maxFreq = max(maxFreq, state[c])

		if k+maxFreq < end-start+1 {
			state[string(s[start])] -= 1
			start++
		}

		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}
