package main

// Given an array of integers representing the value of cards, write a function to calculate the
// maximum score you can achieve by selecting exactly k cards from either the beginning or the end
// of the array.

// For example, if k = 3, then you have the option to select:
// the first three cards,
// the last three cards,
// the first card and the last two cards
// the first two cards and the last card

// example:
// imput: cards = [2,11,4,5,3,9,2]
// 	   k = 3
// output: 17
// Explanation:
// Selecting the first three cards from the beginning (2 + 11 + 4) gives a total of 17.
// Selecting the last three cards (3 + 9 + 2) gives a total of 14.
// Selecting the first card and the last two cards (2 + 9 + 2) gives a total of 13.
// Selecting the first two cards and the last card (2 + 11 + 2) gives a total of 15.
// So the maximum score is 17.

func maxPointsYouCAnObtainFromCards(cardPoints []int, k int) int {
	totalSum := 0
	for _, v := range cardPoints {
		totalSum += v
	}

	if len(cardPoints) >= k {
		return totalSum
	}

	state, start, maxCount := 0, 0, 0

	for end := range len(cardPoints) {
		state += cardPoints[end]

		if end-start+1 == len(cardPoints)-k {
			maxCount = max(maxCount, totalSum-state)
			state -= cardPoints[start]
			start++
		}
	}

	return maxCount
}
