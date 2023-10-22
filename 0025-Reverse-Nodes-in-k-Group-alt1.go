/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k <= 0 {return head}
    var prevGroupEnding, next, curr, prev, tail *ListNode
    for curr = head; curr != nil; {
        if !walkable(curr, k) {break}
        tail = curr
        for i := 0; i < k; i++ {
            next = curr.Next
            curr.Next = prev
            prev = curr
            curr = next
        }
        if prevGroupEnding == nil {
            head = prev
        } else {
            prevGroupEnding.Next = prev
        }
        tail.Next = curr
        prevGroupEnding = tail
        prev = nil
    }
    return head
}

func walkable(curr *ListNode, k int) bool {
    i := 0
    for curr != nil && i < k {
        curr = curr.Next
        i++
    }
    return i == k
}
