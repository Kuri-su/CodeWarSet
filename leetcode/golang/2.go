package golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := l1
	nextPlus := 0
	lastL1 := l1
	for {
		if l1 != nil && l2 != nil {
			l1.Val += l2.Val + nextPlus
			nextPlus = 0

			if l1.Val >= 10 {
				nextPlus = l1.Val / 10
				l1.Val %= 10
			}
		} else if l1 != nil && l2 == nil {
			l1.Val += nextPlus
			nextPlus = 0
			if l1.Val >= 10 {
				nextPlus = l1.Val / 10
				l1.Val %= 10
			}
			lastL1 = l1
			l1 = l1.Next
			continue
		} else if l1 == nil && l2 != nil {
			l1 = &ListNode{
				Val:  0,
				Next: nil,
			}
			l1.Val += l2.Val + nextPlus
			nextPlus = 0

			if l1.Val >= 10 {
				nextPlus = l1.Val / 10
				l1.Val %= 10
			}
			lastL1.Next = l1

		} else if l1 == nil && l2 == nil && nextPlus != 0 {
			l1 = &ListNode{
				Val:  nextPlus,
				Next: nil,
			}
			nextPlus = 0
			lastL1.Next = l1
		} else {
			break
		}

		lastL1 = l1
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return root
}
