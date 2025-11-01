// Package stringutil provides string manipulation utilities for common string operations.
package stringutil

import "strings"

// Reverse returns the input string with its characters in reverse order.
// Properly handles multi-byte UTF-8 characters (runes).
// Example: Reverse("Hello") returns "olleH"
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ToUpper converts the entire string to uppercase.
// Example: ToUpper("hello") returns "HELLO"
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// IsPalindrome checks if a string is a palindrome (reads the same forwards and backwards).
// Comparison is case-insensitive and ignores spaces.
// Example: IsPalindrome("A man a plan a canal Panama") returns true
func IsPalindrome(s string) bool {
	normalized := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return normalized == Reverse(normalized)
}

// helper is unexported (private) - only usable within this package.
// Removes leading and trailing whitespace from a string.
func helper(s string) string {
	return strings.TrimSpace(s)
}
