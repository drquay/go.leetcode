package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := fmt.Sprint(x)
	reverseStr := reverse(original)

	if original == reverseStr {
		return true
	}
	return false
}

func reverse(str string) string {
	var result string
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func main() {

	// Input: x = 121
	// Output: true

	// Input: x = -121
	// Output: false
	// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

	println(isPalindrome(121))
	println(isPalindrome(1221))
	println(isPalindrome(-121))
	println(isPalindrome(10))
}
