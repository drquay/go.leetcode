package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	root := make([]int, 0)
	for true {
		minIndex := -1
		minValue := math.MaxInt64
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil && minValue > lists[i].Val {
				minValue = lists[i].Val
				minIndex = i
			}
		}
		if minIndex != -1 {
			lists[minIndex] = lists[minIndex].Next
			root = append(root, minValue)
		} else {
			break
		}
	}
	return convertToListNode(root)
}

func convertToListNode(val []int) *ListNode {

	var root *ListNode = nil
	for i := len(val) - 1; i >= 0; i-- {
		node := &ListNode{Val: val[i], Next: root}
		root = node
	}
	return root
}

func main() {

	// Input: lists = [[1,4,5],[1,3,4],[2,6]]
	// Output: [1,1,2,3,4,4,5,6]
	// Explanation: The linked-lists are:
	// [
	// 1->4->5,
	// 1->3->4,
	// 2->6
	// ]
	// merging them into one sorted list:
	// 1->1->2->3->4->4->5->6

	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}}

	result := mergeKLists([]*ListNode{list1, list2, list3})
	cur := result
	for cur != nil {
		print(cur.Val)
		cur = cur.Next
		if cur != nil {
			print("->")
		}
	}
	println()
}
