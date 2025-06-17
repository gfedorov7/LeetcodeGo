package tasks_from_leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	remained := 0
	var head *ListNode
	var current *ListNode

	for l1 != nil && l2 != nil {
		valueL1 := l1.Val
		valueL2 := l2.Val

		l1 = l1.Next
		l2 = l2.Next

		finalValue := valueL1 + valueL2
		if remained > 0 {
			finalValue += remained
			remained = 0
		}
		if finalValue >= 10 {
			remained = finalValue - finalValue%10
			if remained%10 == 0 {
				remained = 1
			}
			finalValue = finalValue % 10
		}

		newNode := &ListNode{Val: finalValue}

		if head == nil {
			head = newNode
			current = newNode
		} else {
			current.Next = newNode
			current = current.Next
		}
	}

	if l2 != nil {
		l1 = l2
	}

	for l1 != nil {
		finalValue := l1.Val
		if remained > 0 {
			finalValue += remained
			remained = 0
		}
		if finalValue >= 10 {
			remained = finalValue - finalValue%10
			if remained%10 == 0 {
				remained = 1
			}
			finalValue = finalValue % 10
		}

		newNode := &ListNode{Val: finalValue}
		l1 = l1.Next

		current.Next = newNode
		current = current.Next
	}

	if remained > 0 {
		current.Next = &ListNode{Val: remained}
		current = current.Next
	}

	return head
}
