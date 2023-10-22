/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(node *ListNode) *ListNode {
    if node == nil || node.Next == nil {
        return node
    }
    prev := node.Next
    head := reverseList(node.Next)
    prev.Next = node
    node.Next = nil
    return head
}
