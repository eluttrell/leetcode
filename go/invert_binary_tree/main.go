package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) TreeToArr() []int {
	tArr := []int{}
	buildTreeArr(t, 0, &tArr)
	return tArr
}

func buildTreeArr(t *TreeNode, idx int, tArr *[]int) {
	tArrLen := len(*tArr)

	if idx >= tArrLen {
		*tArr = append((*tArr), make([]int, idx-(tArrLen-1))...)
	}

	if t.Left != nil {
		buildTreeArr(t.Left, (idx*2)+1, tArr)
	}
	if t.Right != nil {
		buildTreeArr(t.Right, (idx*2)+2, tArr)
	}

	(*tArr)[idx] = t.Val

	return
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		if root.Left == nil {
			root.Left = invertTree(root.Right)
			root.Right = nil
		} else if root.Right == nil {
			root.Right = invertTree(root.Left)
			root.Left = nil
		} else {
			temp := root.Left
			root.Left = invertTree(root.Right)
			root.Right = invertTree(temp)
		}
	}

	return root
}

func buildTree(arr *[]int, start int) *TreeNode {
	if len(*arr) == 0 {
		return &TreeNode{}
	}

	leftIdx := (start * 2) + 1
	rightIdx := (start * 2) + 2

	root := TreeNode{Val: (*arr)[start]}

	if leftIdx < len(*arr) {
		root.Left = buildTree(arr, leftIdx)
	}
	if rightIdx < len(*arr) {
		root.Right = buildTree(arr, rightIdx)
	}

	return &root
}

func main() {}
