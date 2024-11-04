package main

import "fmt"

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
	arr := []int{1, 2, 5, 2, 8, 1, 5}
	num := 3
	result := maxSubarraySum(arr, num)
	fmt.Println("Maximum sum of subarray of size", num, "is:", result)
}
