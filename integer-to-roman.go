//https://leetcode.com/problems/integer-to-roman/
package main

func intToRoman(num int) string {

	res := string("")
	units := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(units); i++ {
		unit := units[i]
		if num >= unit {
			n := num / unit
			for j := 0; j < n; j++ {
				res += strs[i]
			}
			num %= unit
		}
	}

	return res
}

func main() {

	// Symbol       Value
	// I             1
	// V             5
	// X             10
	// L             50
	// C             100
	// D             500
	// M             1000
	// Input: num = 3
	// Output: "III"

	println(intToRoman(3))
	println(intToRoman(4))
	println(intToRoman(9))
	println(intToRoman(58))
	println(intToRoman(19994))
}
