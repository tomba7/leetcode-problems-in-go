/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    var firstDummy, secondDummy ListNode
    first, second := &firstDummy, &secondDummy
    for curr := head; curr != nil; curr = curr.Next {
        if curr.Val < x {
            first.Next = curr
            first = first.Next
        } else {
            second.Next = curr
            second = second.Next
        }
    }
    second.Next = nil
    first.Next = secondDummy.Next
    return firstDummy.Next
}
