package main

// Given a list of intervals intervals and an interval newInterval, write a function to insert
// newInterval into a list of existing, non-overlapping, and sorted intervals based on their
// starting points. The function should ensure that after the new interval is added, the list
// remains sorted without any overlapping intervals, merging them if needed.
// example:
// input: intervals = [[1,3],[6,9]]
//
//	newInterval = [2,5]
//
// output: [[1,5],[6,9]]
// Explanation: The new interval [2,5] overlaps with [1,3], so they are merged into [1,5].
func insertIntervals(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	n := len(intervals)

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
