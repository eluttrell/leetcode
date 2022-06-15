package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	// Strip non-alphanumeric and split to array of characters
	str := strings.Split(reg.ReplaceAllString(strings.ToLower(s), ""), "")

	length := len(str)
	ptr1 := (length / 2) - 1
	ptr2 := length / 2

	if length%2 != 0 {
		ptr2++
	}

	for ptr1 >= 0 && ptr2 < length {
		if str[ptr1] == str[ptr2] {
			ptr1--
			ptr2++
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "A man, a plan, a canal: Panama!"
	s1 := "race a car"
	s2 := "ab"

	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindrome(s1))
	fmt.Println(isPalindrome(s2))
}
