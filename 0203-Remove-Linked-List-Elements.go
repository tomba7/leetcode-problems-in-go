/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := ListNode{Next:head}
    prev, curr := &dummy, head
    for curr != nil {
        next := curr.Next
        if curr.Val == val {
            prev.Next = next
        } else {
            prev = prev.Next
        }
        curr = next
    }
    return dummy.Next
}
