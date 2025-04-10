package main

import "sort"

// Write a function to consolidate overlapping intervals within a given array intervals, where each
// interval intervals[i] consists of a start time starti and an end time endi. The function should
// return an array of the merged intervals so that no two intervals overlap and all the intervals
// collectively cover all the time ranges in the original input.
// examples:
// input: intervals = [[3,5],[1,4],[7,9],[6,8]]
// output: [[1,5],[6,9]]
func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	if len(intervals) == 1 {
		return intervals
	}

	rslt := [][]int{intervals[0]}
	i := 1

	for i < len(intervals) {
		if intervals[i][0] <= rslt[len(rslt)-1][1] {
			rslt[len(rslt)-1][1] = max(rslt[len(rslt)-1][1], intervals[i][1])
		} else {
			rslt = append(rslt, intervals[i])
		}
		i++
	}

	return rslt
}
