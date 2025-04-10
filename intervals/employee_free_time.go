package main

import (
	"log"
	"sort"
)

// Write a function to find the common free time for all employees from a list called schedule. Each
// employee's schedule is represented by a list of non-overlapping intervals sorted by start times.
// The function should return a list of finite, non-zero length intervals where all employees are
// free, also sorted in order.
// example:
// input: schedule = [[[2,4],[7,10]],[[1,5]],[[6,9]]]
// output: [(5,6)]
// Explanation: The three employees collectively have only one common free time interval, which is
// from 5 to 6.
func employeeFreeTime(schedule [][][]int) [][]int {
	flattened := [][]int{}
	for i := range schedule {
		flattened = append(flattened, schedule[i]...)
	}
	sort.Slice(flattened, func(i, j int) bool {
		return flattened[i][0] < flattened[j][0]
	})

	log.Println("flattened", flattened)

	merged := [][]int{
		flattened[0],
	}

	for i := 1; i < len(flattened); i++ {
		if flattened[i][0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], flattened[i][1])
		} else {
			merged = append(merged, flattened[i])
		}
	}

	freeTime := [][]int{}
	for i := 1; i < len(merged); i++ {
		freeTime = append(freeTime, []int{merged[i-1][1], merged[i][0]})
	}

	return freeTime
}
