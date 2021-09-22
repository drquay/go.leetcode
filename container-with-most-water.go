// https://leetcode.com/problems/container-with-most-water/
package main

func maxArea(height []int) int {
	n := len(height)
	maxArea := 0
	i := 0
	j := n - 1
	for i < j {
		var value int
		if value = height[i]; value > height[j] {
			value = height[j]
		}
		area := (j - i) * value
		if area > maxArea {
			maxArea = area
		}

		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func main() {
	// Input: height = [1,8,6,2,5,4,8,3,7]
	// Output: 49
	// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
	// In this case, the max area of water (blue section) the container can contain is 49.

	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))     // 17
}
