//https://leetcode.com/problems/longest-palindromic-substring/
package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	n := len(s)
	start := 0
	maxLenght := 1
	var left, right int
	for i := 1; i < n; i++ {
		left = i - 1
		right = i
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		left++
		right--
		if s[left] == s[right] && maxLenght < right-left+1 {
			maxLenght = right - left + 1
			start = left
		}
		left = i - 1
		right = i + 1
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		left++
		right--
		if s[left] == s[right] && maxLenght < right-left+1 {
			maxLenght = right - left + 1
			start = left
		}
	}
	return s[start : start+maxLenght]
}

func main() {
	// Input: s = "babad"
	// Output: "bab"
	// Note: "aba" is also a valid answer.

	// Input: s = "cbbd"
	// Output: "bb"

	println(longestPalindrome("babad"))
	println(longestPalindrome("cbbd"))
	println(longestPalindrome("c"))
}
