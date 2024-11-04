package main

import "fmt"

// count down with recursion.
func countDown(num int) {
	if num <= 0 {
		fmt.Println("Done")

		return
	}

	fmt.Println(num)
	countDown(num - 1)
}

// calculates the sum of a certain range.
func sumRange(num int) int {
	if num == 1 {
		return 1
	}

	return num + sumRange(num-1)
}

// calculates a factorial of a number
func factorial(num int) int {
	if num == 2 {
		return 2
	}

	return num * factorial(num-1)
}

func main() {
	// countDown(5)
	// fmt.Println(sumRange(4))
	fmt.Println(factorial(5))
}
