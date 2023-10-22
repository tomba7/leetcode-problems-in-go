/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    slow, fast := head, head
    var prev *ListNode

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        next := slow.Next
        slow.Next = prev
        prev = slow
        slow = next
    }

    first, second := prev, slow
    if fast != nil { second = second.Next }

    for second != nil {
        if first.Val != second.Val {
            return false
        }
        first, second = first.Next, second.Next
    }
    return true
}
