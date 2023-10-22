/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {return nil}
    left, right := head, head
    var prev *ListNode
    for right != nil {
        if n > 0 {
            n--
        } else {
            prev = left
            left left.Next
        }
        right = right.Next
    }
    if prev == nil {
        head = head.Next
    } else {
        prev.Next = left.Next
    }
    return head
}
