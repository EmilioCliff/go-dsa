package main

import (
	"log"
)

func maxSubarraySum(arr []int, num int) int {
	if len(arr) < num || num <= 0 {
		return 0
	}

	maxSum := 0
	for i := 0; i < num; i++ {
		maxSum += arr[i]
	}

	currentSum := maxSum
	for i := num; i < len(arr); i++ {
		currentSum = currentSum - arr[i-num] + arr[i]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	arr1 := []int{1, 2, 5, 2, 8, 1, 5}
	log.Println(maxSubarraySum(arr1, 3))

	arr2 := []int{3, 3, 2, 1, 2, 1, 0}
	arr3 := []int{0, 1, 2, 2}
	log.Println(fruitIntoBaskets(arr2))
	log.Println(fruitIntoBaskets(arr3))

	arr4 := "abcabcbb"
	arr5 := "pwwkew"
	arr6 := "dvdf"
	arr7 := "tmmzuxt"
	log.Println(longestSubstringWithoutRepeatingChars(arr4))
	log.Println(longestSubstringWithoutRepeatingChars(arr5))
	log.Println(longestSubstringWithoutRepeatingChars(arr6))
	log.Println(longestSubstringWithoutRepeatingCharsOpt(arr7))

	arr8 := "AABABBA"
	arr9 := "BBABCCDD"
	log.Println(longestRepeatingCharReplacement(arr8, 1))
	log.Println(longestRepeatingCharReplacement(arr9, 2))

	arr10 := []int{2, 1, 5, 1, 3, 2}
	log.Println(maxSumOfSubarrayWithSizeK(arr10, 3))

	arr11 := []int{2, 11, 4, 5, 3, 9, 2}
	arr12 := []int{9, 7, 7, 9, 7, 7, 9}
	log.Println(maxPointsYouCAnObtainFromCards(arr11, 3))
	log.Println(maxPointsYouCAnObtainFromCards(arr12, 7))

	arr13 := []int{3, 2, 2, 3, 4, 6, 7, 7, -1}
	log.Println(maxSumOfDistinctSubarraysLenghtK(arr13, 4))
}
