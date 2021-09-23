package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	var arr []int
	for i := 0; i < n; i++ {
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			arr = append(arr, sum)
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	min := absDiffInt(target, arr[0])
	index := 0
	for i := 0; i < len(arr); i++ {
		val := absDiffInt(target, arr[i])
		if min > val {
			min = val
			index = i
		}
	}
	return arr[index]
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {

	// Input: nums = [-1,2,1,-4], target = 1
	// Output: 2
	// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

	// Input: nums = [0,0,0], target = 1
	// Output: 0

	// Input: nums = [1,1,1,0], target = -100
	// Output: 2

	// println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	// println(threeSumClosest([]int{0, 0, 0}, 0))
	println(threeSumClosest([]int{1, 1, 1, 0}, -100))
}
