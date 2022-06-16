package main

import (
	"reflect"
	"strings"
)

func isAnagrams(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	sFreq := make(map[string]int)
	tFreq := make(map[string]int)

	sStr := strings.Split(s, "")
	tStr := strings.Split(t, "")

	for i, v := range sStr {
		if _, ok := sFreq[v]; !ok {
			sFreq[v] = 1
		} else {
			sFreq[v]++
		}
		if _, ok := tFreq[tStr[i]]; !ok {
			tFreq[tStr[i]] = 1
		} else {
			tFreq[tStr[i]]++
		}
	}

	return reflect.DeepEqual(sFreq, tFreq)
}

func main() {}
