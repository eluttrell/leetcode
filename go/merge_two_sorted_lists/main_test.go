package main

import "testing"

func TestMergeTwoLists(t *testing.T) {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	list1 := buildList(&arr1)
	list2 := buildList(&arr2)

	expected := []int{1, 1, 2, 3, 4, 4}

	merged := mergeTwoLists(list1, list2)

	mergedArr := []int{0}

	merged.ListToArr(&mergedArr)

	match := true

	for i, v := range mergedArr {
		if i > len(mergedArr) || v != mergedArr[i] {
			match = false
		}
	}

	if !match {
		t.Errorf("Merge failed, expected %v, got %v", expected, mergedArr)
	}
}
