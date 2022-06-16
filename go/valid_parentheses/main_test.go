package main

import "testing"

type TestCase struct {
	Str   string
	valid bool
}

func TestIsPalindrome(t *testing.T) {
	tests := []TestCase{
		{Str: "{{{}}}", valid: true},
		{Str: "{}[](){[()]}", valid: true},
		{Str: "{]", valid: false},
		{Str: "}[", valid: false},
	}

	for _, v := range tests {
		result := isValid(v.Str)

		if result != v.valid {
			t.Errorf("Error testing string, expected Palindrome response to be %v, got %v", v.valid, result)
		}
	}
}
