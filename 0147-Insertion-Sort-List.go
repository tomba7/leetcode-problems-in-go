/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    var dummyHead ListNode
    for node := head; node != nil; {
        next := node.Next
        node.Next = nil
        prev := &dummyHead
        for prev.Next != nil && prev.Next.Val < node.Val {
            prev = prev.Next
        }
        node.Next = prev.Next
        prev.Next = node
        node = next
    }
    return dummyHead.Next
}
