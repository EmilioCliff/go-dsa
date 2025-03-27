package main

import (
	"log"
)

func main() {
	arr1 := []int{1, 2, 1}
	log.Println(containerWithMostWater(arr1))

	arr2 := []int{-1, 0, 1, 2, -1, -1}
	log.Println(threeSum(arr2))

	arr3 := []int{11, 4, 9, 6, 15, 18}
	log.Println(triangleNumbers(arr3))

	arr4 := []int{2, 0, 4, 0, 9}
	log.Println(moveZeroes(arr4))

	arr5 := []int{2, 1, 2, 0, 1, 0, 1, 0, 1}
	log.Println(sortColors(arr5))

	arr6 := []int{3, 4, 1, 2, 2, 5, 1, 0, 2}          // 10
	arr7 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1} // 6
	arr8 := []int{5, 4, 1, 2}                         // 1
	log.Println(trappingRainWaterNotAll(arr6))
	log.Println(trappingRainWaterNotAll(arr7))
	log.Println(trappingRainWaterNotAll(arr8))

	log.Println(trappingRainWater(arr6))
	log.Println(trappingRainWater(arr7))
	log.Println(trappingRainWater(arr8))
}
