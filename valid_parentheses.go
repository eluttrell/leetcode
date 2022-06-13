package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	temp := []string{}
	str := strings.Split(s, "")

	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i < len(str); i++ {
		if str[i] == "(" || str[i] == "{" || str[i] == "[" {
			temp = append(temp, str[i])
		} else {
			if len(temp) <= 0 {
				return false
			}
			last := temp[len(temp)-1]
			temp = append([]string{}, temp[:len(temp)-1]...)
			switch str[i] {
			case ")":
				if last != "(" {
					return false
				}
			case "}":
				if last != "{" {
					return false
				}
			case "]":
				if last != "[" {
					return false
				}
			}
		}
	}

	return len(temp) == 0
}

func main() {
	s := "{{{}}}"        // Good
	s1 := "{}[](){[()]}" // Good
	s2 := "{]"           // Bad
	s3 := "}["           // Bad
	fmt.Println(isValid(s) == true)
	fmt.Println(isValid(s1) == true)
	fmt.Println(isValid(s2) == false)
	fmt.Println(isValid(s3) == false)
}
