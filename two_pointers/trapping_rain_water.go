package main

// Write a function to calculate the total amount of water trapped between bars on an elevation map,
// where each bar's width is 1. The input is given as an array of n non-negative integers height
// representing the height of each bar.

// example:
// input: height = [3, 4, 1, 2, 2, 5, 1, 0, 2]
// output: 10

func trappingRainWaterNotAll(height []int) int {
	count := 0
	left := 0

	for left < len(height)-1 {
		right := left + 1
		highest := height[right]
		for right < len(height)-1 && height[right] < height[left] {
			right++
			highest = max(highest, height[right])
		}

		smallWall := min(height[left], highest)
		reached := right

		for left < right {
			right--
			if left == right {
				break
			}
			count += smallWall - height[right]
		}
		left = reached
	}

	return count
}

func trappingRainWater(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	count := 0

	for left+1 < right {
		if leftMax > rightMax {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				count += rightMax - height[right]
			}
		} else {
			left++
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				count += leftMax - height[left]
			}
		}
	}

	return count
}
