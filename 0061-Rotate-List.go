/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {return head}
    var left, right, prev, tail *ListNode = head, head, nil, nil
    k %= getLength(head)
    if k == 0 {return head}
    for i := 0; right != nil && i < k; i++ {
        right = right.Next
    }
    for right != nil {
        prev, tail = left, right
        left, right = left.Next, right.Next
    }
    prev.Next = nil
    tail.Next = head
    return left
}

func getLength(head *ListNode) (length int) {
    for curr := head; curr != nil; curr = curr.Next {
        length++
    }
    return
}
