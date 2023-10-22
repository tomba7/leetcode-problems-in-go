/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var prev, next *ListNode
    curr := head
    for curr != nil {
        next = curr.Next
        if prev != nil && prev.Val == curr.Val {
            prev.Next = next
        } else {
            prev = curr
        }
        curr = next
    }
    return head
}
