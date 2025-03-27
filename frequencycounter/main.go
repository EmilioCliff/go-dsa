package main

import (
	"log"
)

// same determines if two arrays have the same elements, with the same multiplicity,
// even if they are in a different order.
func same(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	arr1FC := map[int]int{}
	arr2FC := map[int]int{}

	for _, n := range arr1 {
		arr1FC[n]++
	}

	for _, n := range arr2 {
		arr1FC[n]++
	}

	for k, v := range arr1FC {
		if _, ok := arr2FC[k*k]; !ok {
			return false
		}

		if v2 := arr2FC[k*k]; v != v2 {
			return false
		}
	}

	return true
}

// sameString determines if two strings are anagrams of each other.
// It checks if both strings contain the same characters with the same frequency,
// irrespective of their order. Returns true if they are anagrams, otherwise false.
func sameString(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	lookup := map[string]int{}

	for i := 0; i < len(str1); i++ {
		char := string(str1[i])
		lookup[char]++
	}

	for i := 0; i < len(str2); i++ {
		char := string(str2[i])
		if val, ok := lookup[char]; !ok || val == 0 {
			return false
		} else {
			lookup[char] -= 1
		}
	}

	return true
}

func main() {
	arr1 := []int{1, 2, 3, 2, 5}
	arr2 := []int{9, 1, 9, 9, 25}
	log.Println(same(arr1, arr2))

	str := map[int][]string{
		1: {"", ""},
		2: {"aaz", "zza"},
		3: {"anagram", "nagaram"},
		4: {"rat", "car"},
		5: {"awesome", "awesom"},
		6: {"qwerty", "qeywrt"},
		7: {"texttwisttime", "timetwisttext"},
	}

	for k, v := range str {
		log.Printf("%d: %v", k, sameString(v[0], v[1]))
	}
}
