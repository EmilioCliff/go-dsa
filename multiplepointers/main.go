package main

import (
	"errors"
	"log"
)

// finds the numbers in an array that adds the sum to zero
// uses two pointers and a sorted array where if sum < 0 left pointer moves forward and vice versa
func sumZero(arr []int) ([]int, error) {
	s := 0
	e := len(arr) - 1

	for s < e {
		sum := arr[s] + arr[e]
		if sum < 0 {
			s++

			continue
		}

		if sum == 0 {
			return []int{arr[s], arr[e]}, nil
		}

		e--
	}

	return nil, errors.New("undefined")
}

// countUniqueValues counts the number of unique values in a sorted integer array.
// It returns the count of unique values. An empty array returns 0.
func countUniqueValues(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	l := 0
	result := 1

	for i, v := range arr {
		if v == arr[l] {
			continue
		}

		l = i
		result += 1
	}

	return result
}

// countUniqueValues2 counts the number of unique values in a sorted integer array.
// It returns the count of unique values. An empty array returns 0.
//
// The function uses an in-place algorithm, so it modifies the input array.
func countUniqueValues2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i]
		}
	}

	return j + 1
}

func main() {
	// arr1 := []int{-3, -2, -1, 0, 1, 2, 3}
	// arr2 := []int{-2, 0, 1, 3}
	// arr3 := []int{1, 2, 3}
	// arr4 := []int{-4, -3, -2, -1, 0, 1, 2, 5}
	// samples := [][]int{arr1, arr2, arr3, arr4}

	// for i, v := range samples {
	// 	rsp, err := sumZero(v)
	// 	log.Printf("%d: result=%v err=%v", i, rsp, err)
	// }

	arr1 := []int{1, 1, 1, 1, 1, 2}
	arr2 := []int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}
	arr3 := []int{}
	arr4 := []int{-2, -1, -1, 0, 1}
	arr5 := []int{0, 0}
	samples := [][]int{arr1, arr2, arr3, arr4, arr5}

	for i, v := range samples {
		rsp := countUniqueValues2(v)
		log.Printf("%d: result=%v", i, rsp)
	}
}
