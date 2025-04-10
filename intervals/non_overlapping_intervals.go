package main

import "sort"

// Write a function to return the minimum number of intervals that must be removed from a given
// array intervals, where intervals[i] consists of a starting point starti and an ending point endi,
// to ensure that the remaining intervals do not overlap.
// examples:
// input: intervals = [[1,3],[5,8],[4,10],[11,13]]
// output: 1
func nonOverlappingIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	end := intervals[0][1]
	count := 1

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}

	return len(intervals) - count
}
