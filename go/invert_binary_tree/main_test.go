package main

import "testing"

func TestInvertTree(t *testing.T) {
	treeArr := []int{4, 2, 7, 1, 3, 6, 9, 8, 2, 1, 1, 2, 6, 3, 4}
	expected := []int{4, 7, 2, 9, 6, 3, 1, 4, 3, 6, 2, 1, 1, 2, 8}

	root := buildTree(&treeArr, 0)

	rootArr := root.TreeToArr()

	treeBuiltCorrectly := true

	for i, v := range rootArr {
		if i > len(treeArr) || treeArr[i] != v {
			treeBuiltCorrectly = false
		}
	}

	if !treeBuiltCorrectly {
		t.Errorf("TreeToArr method not working, expected %v, got %v", treeArr, rootArr)
	}

	invertedTree := invertTree(root)

	invertedTreeArr := invertedTree.TreeToArr()

	treeInvertedCorrectly := true

	for i, v := range invertedTreeArr {
		if i > len(treeArr) || expected[i] != v {
			treeInvertedCorrectly = false
		}
	}

	if !treeInvertedCorrectly {
		t.Errorf("invertTree failed, expected %v, go %v", expected, invertedTreeArr)
	}
}
