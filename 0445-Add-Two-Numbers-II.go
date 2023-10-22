/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
    m, n := getLength(l1), getLength(l2)
    if n > m {
        // Make sure l1 is the longer list
        l1, l2 = l2, l1
        m, n = n, m
    }
    next, carry := addTwoNumbersHelper(l1, l2, m - n)
    if carry != 0 {
        return &ListNode{Val: carry, Next: next}
    }
    return next
}

func addTwoNumbersHelper(l1, l2 *ListNode, diff int) (*ListNode, int) {
    if l1 == nil && l2 == nil {return nil, 0}
    var next *ListNode
    var sum int
    if diff > 0 {
        next, sum = addTwoNumbersHelper(l1.Next, l2, diff - 1)
        sum += l1.Val
    } else {
        next, sum = addTwoNumbersHelper(l1.Next, l2.Next, diff)
        sum += l1.Val + l2.Val
    }
    return &ListNode{Val: sum % 10, Next: next}, sum / 10
}

func getLength(head *ListNode) (length int) {
    for curr := head; curr != nil; curr = curr.Next {length++}
    return
}
