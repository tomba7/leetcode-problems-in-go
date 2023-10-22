/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var oddPrev *ListNode
    odd, even, evenHead := head, head.Next, head.Next
    for odd != nil && even != nil {
        oddPrev = odd
        odd.Next = even.Next
        odd = odd.Next
        if odd != nil {
            even.Next = odd.Next
            even = even.Next
        }
    }
    if odd != nil {
        odd.Next = evenHead
    } else {
        oddPrev.Next = evenHead
    }
    return head
}
