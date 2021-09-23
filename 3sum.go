//https://leetcode.com/problems/3sum/
package main

import "sort"

func threeSum(nums []int) [][]int {

	n := len(nums)
	if n < 2 {
		return [][]int{}
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

func main() {
	// Input: nums = [-1,0,1,2,-1,-4]
	// Output: [[-1,-1,2],[-1,0,1]]
	// nums[i] + nums[j] + nums[k] == 0

	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for i := 0; i < len(res); i++ {
		print("[ ")
		for j := 0; j < len(res[i]); j++ {
			print(res[i][j], " ")
		}
		println("]")
	}
}
