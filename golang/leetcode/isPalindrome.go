package leetcode

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	isValid := func(r rune) bool {
		return unicode.IsDigit(r) ||unicode.IsLetter(r)
	}
	l := len(s)
	i, j := 0, l - 1
	var a, b rune
	for i < j {
		a, b = rune(s[i]), rune(s[j])
		if !isValid(a) && !isValid(b) {
			i++
			j--
		} else if !isValid(a) {
			i++
		} else if !isValid(b) {
			j--
		} else if unicode.ToLower(a) == unicode.ToLower(b) {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}