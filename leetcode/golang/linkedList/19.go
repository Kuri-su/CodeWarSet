package linkedList

/*
   双指针解法

   解题思想:


   时间复杂度: O(n)
   空间复杂度: O(1)

   过程:
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	p2 := head
	len := 0

	m := n
	for {
		len++

		if p1.Next == nil {
			break
		} else {
			p1 = p1.Next
		}

		if m > 0 {
			m--
		} else {
			p2 = p2.Next
		}

	}
	if len <= 1 {
		return nil
	}
	if len == n {
		return head.Next
	}

	if n == 1 {
		p2.Next = nil
	} else {
		p2.Next = p2.Next.Next
	}

	return head
}
