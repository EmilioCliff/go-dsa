package main

import "log"

func main() {
	log.Println(testOne("substring"))
	log.Println(testTwo("BBABCCDD", 2))
}

func testOne(s string) int {
	state := map[string]int{}
	start := 0
	maxLength := 0

	for end := range s {
		c := string(s[end])
		state[c] += 1

		for state[c] > 1 {
			state[string(s[start])] -= 1
			start++
		}

		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

func testTwo(s string, k int) int {
	state := map[string]int{}
	start, maxFreq, maxCount := 0, 0, 0

	for end := range s {
		c := string(s[end])
		state[c] += 1
		maxFreq = max(maxFreq, state[c])

		if k+maxFreq < end-start+1 {
			state[string(s[start])] -= 1
			start++
		}

		maxCount = max(maxCount, end-start+1)
	}

	return maxCount
}
