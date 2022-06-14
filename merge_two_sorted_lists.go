package main

import "fmt"

// Singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Singly linked list Print method
func (l *ListNode) Print() {
	ptr := l
	for ptr != nil {
		fmt.Println("Node: ", ptr.Val)
		ptr = ptr.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	list3 := ListNode{}
	cur := &list3

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Val = list1.Val
			cur.Next = &ListNode{}
			cur = cur.Next
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			cur.Val = list2.Val
			cur.Next = &ListNode{}
			cur = cur.Next
			list2 = list2.Next
		}

		if list1 == nil {
			for list2 != nil {
				cur.Val = list2.Val
				if list2.Next != nil {
					cur.Next = &ListNode{}
					cur = cur.Next
				}
				list2 = list2.Next
			}
		} else if list2 == nil {
			for list1 != nil {
				cur.Val = list1.Val
				if list1.Next != nil {
					cur.Next = &ListNode{}
					cur = cur.Next
				}
				list1 = list1.Next
			}
		}
	}

	return &list3
}

func main() {
	// Build linked lists from slices
	x := ListNode{}
	y := ListNode{}

	xl := &x
	yl := &y

	lx := []int{1, 2, 4}
	ly := []int{1, 3, 4}

	for i := 0; i < 3; i++ {
		xl.Val = lx[i]
		yl.Val = ly[i]
		if i < 2 {
			xl.Next = &ListNode{}
			yl.Next = &ListNode{}
			xl = xl.Next
			yl = yl.Next
		}
	}

	mergeTwoLists(&x, &y).Print()
}
