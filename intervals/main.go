package main

import (
	"log"
)

func main() {
	// intervalsOne := [][]int{
	// 	{1, 5},
	// 	{3, 9},
	// 	{6, 8},
	// }
	// intervalsTwo := [][]int{
	// 	{10, 12},
	// 	{6, 9},
	// 	{13, 15},
	// }
	// log.Println(canAttendMettings(intervalsOne))
	// log.Println(canAttendMettings(intervalsTwo))

	intervalThree := [][]int{
		{1, 3},
		{6, 9},
	}
	intervalFour := [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	log.Println(insertIntervals(intervalThree, []int{2, 5}))
	log.Println(insertIntervals(intervalFour, []int{4, 8}))
	log.Println(nonOverlappingIntervals([][]int{
		{1, 3},
		{5, 8},
		{4, 10},
		{11, 13},
	}))
	log.Println(mergeIntervals([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))

	log.Println(employeeFreeTime([][][]int{
		{
			{2, 4},
			{7, 10},
		},
		{
			{1, 5},
		},
		{
			{6, 9},
		},
	}))
}
