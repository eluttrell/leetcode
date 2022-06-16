package main

// Singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Singly linked list Print method
func (l *ListNode) ListToArr(arr *[]int) {
	ptr := l
	counter := 0
	for ptr != nil {
		length := len(*arr)
		if length <= counter {
			(*arr) = append((*arr), make([]int, counter-(length-1))...)
		}
		(*arr)[counter] = ptr.Val
		counter++
		ptr = ptr.Next
	}
}

func buildList(arr *[]int) *ListNode {
	length := len(*arr)
	head := ListNode{}
	l := &head
	for i := 0; i < length; i++ {
		l.Val = (*arr)[i]
		if i < length-1 {
			l.Next = &ListNode{}
			l = l.Next
		}
	}

	return &head
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

func main() {}
