package main

import (
	"sort"
)

// Write a function to check if a person can attend all the meetings scheduled without any time
// conflicts. Given an array intervals, where each element [s1, e1] represents a meeting starting at
// time s1 and ending at time e1, determine if there are any overlapping meetings. If there is no
// overlap between any meetings, return true; otherwise, return false. Note that meetings ending and
// starting at the same time, such as (0,5) and (5,10), do not conflict.
// examples:
// input: intervals = [(1,5),(3,9),(6,8)]
// output: false
// Explanation: The meetings (1,5) and (3,9) overlap.
// input: intervals = [(10,12),(6,9),(13,15)]
// output: true
// Explanation: There are no overlapping meetings, so the person can attend all.
func canAttendMettings(meetings [][]int) bool {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for k := range meetings {
		if k == 0 {
			continue
		}
		if meetings[k][0] < meetings[k-1][1] {
			return false
		}
	}

	return true
}
