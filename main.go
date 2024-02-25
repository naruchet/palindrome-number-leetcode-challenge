package main

import (
	"strconv"
)

func main() {
	isPalindrome(121121)
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	count := len(s) / 2

	for i := 0; i < count; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
