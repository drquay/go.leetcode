//https://leetcode.com/problems/zigzag-conversion/
package main

func convert(s string, rowNums int) string {
	if len(s) < 3 || rowNums < 2 {
		return s
	}
	ret := string("")
	for row := 0; row < rowNums; row++ {
		jump := row
		for jump < len(s) {
			ret += string(s[jump])
			if row != 0 && row+1 != rowNums {
				index := 2*rowNums - 2 - 2*row
				if jump+index < len(s) {
					ret += string(s[index+jump])
				}
			}
			jump += 2*rowNums - 2
		}
	}
	return ret
}

func main() {

	// Input: s = "PAYPALISHIRING", numRows = 3
	// Output: "PAHNAPLSIIGYIR"
	// Explanation:
	// P   A   H   N
	// A P L S I I G
	// Y   I   R

	// Input: s = "PAYPALISHIRING", numRows = 4
	// Output: "PINALSIGYAHRPI"
	// Explanation:
	// P     I    N		=> 2*rowNums - 2
	// A   L S  I G		=> 2*rowNums - 2 - 2*row
	// Y A   H R		=> 2*rowNums - 2 - 2*row
	// P     I			=> 2*rowNums - 2

	println(convert("PAYPALISHIRING", 3))
	println(convert("PAYPALISHIRING", 4))
	println(convert("ABC", 1))
}
