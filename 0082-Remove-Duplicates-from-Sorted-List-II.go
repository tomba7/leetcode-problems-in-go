/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var dummy ListNode
    var prev, curr *ListNode = &dummy, head
    for curr != nil && curr.Next != nil {
        if curr.Val != curr.Next.Val {
            prev.Next = curr
            prev = prev.Next
        } else {
            for curr != nil && curr.Next != nil && curr.Val == curr.Next.Val {
                curr = curr.Next
            }
        }
        curr = curr.Next
    }
    prev.Next = curr
    return dummy.Next
}
