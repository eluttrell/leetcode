package main

import "testing"

type TestCase struct {
	Str        string
	Palindrome bool
}

func TestIsPalindrome(t *testing.T) {
	tests := []TestCase{
		{Str: "A man, a plan, a canal: Panama!", Palindrome: true},
		{Str: "race a car", Palindrome: false},
		{Str: "ab", Palindrome: false},
	}

	for _, v := range tests {
		result := isPalindrome(v.Str)

		if result != v.Palindrome {
			t.Errorf("Error testing string, expected Palindrome response to be %v, got %v", v.Palindrome, result)
		}
	}
}
