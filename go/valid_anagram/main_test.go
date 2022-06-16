package main

import "testing"

func TestIsAnagrams(t *testing.T) {
	result1 := isAnagrams("anagram", "nagaram")
	result2 := isAnagrams("panama", "elijah")

	if !result1 {
		t.Errorf("Expected %v, got %v for result1", true, result1)
	}
	if result2 {
		t.Errorf("Expected %v, got %v for result2", false, result2)
	}
}
