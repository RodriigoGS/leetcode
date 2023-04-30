package add_two_numbers

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}
)

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head       ListNode
		carry, sum int

		current     = &head
		left, right = l1, l2
	)

	for left != nil || right != nil {
		var leftVal, rightVal int

		if left != nil {
			leftVal = left.Val
			left = left.Next
		}

		if right != nil {
			rightVal = right.Val
			right = right.Next
		}

		sum = leftVal + rightVal + carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return head.Next
}
