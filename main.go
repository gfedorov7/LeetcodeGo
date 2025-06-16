package main

import (
	"LeetcodeGo/tasks_from_leetcode"
	"fmt"
)

func main() {
	var head *tasks_from_leetcode.ListNode
	var current *tasks_from_leetcode.ListNode

	for _, val := range []int{3, 4, 2} {
		newNode := &tasks_from_leetcode.ListNode{Val: val}

		if head == nil {
			head = newNode
			current = head
		} else {
			current.Next = newNode
			current = current.Next
		}
	}

	var head2 *tasks_from_leetcode.ListNode
	var current2 *tasks_from_leetcode.ListNode

	for _, val := range []int{4, 6, 5} {
		newNode := &tasks_from_leetcode.ListNode{Val: val}

		if head2 == nil {
			head2 = newNode
			current2 = head2
		} else {
			current2.Next = newNode
			current2 = current2.Next
		}
	}

	for current := head; current != nil; current = current.Next {
		fmt.Println(current.Val)
	}

	for current := head2; current != nil; current = current.Next {
		fmt.Println(current.Val)
	}
}
