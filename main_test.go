package main

import "testing"

func Test_Palindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "is palindrome 121",
			input:    121,
			expected: true,
		},
		{
			name:     "is palindrome 1221",
			input:    1221,
			expected: true,
		},
		{
			name:     "is not palindrome -121",
			input:    -121,
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome(tt.input)
			if res != tt.expected {
				t.Errorf("want %t but got %t", tt.expected, res)
			}
		})
	}
}
