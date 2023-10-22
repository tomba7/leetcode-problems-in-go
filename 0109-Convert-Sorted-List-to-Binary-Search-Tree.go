/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    n, node := 0, head
    for node != nil { n, node = n + 1, node.Next }
    return sortedListToBSTHelper(&head, 0, n - 1)
}

func sortedListToBSTHelper(node **ListNode, lo, hi int) *TreeNode {
    if lo > hi { return nil }
    mid := lo + (hi - lo)/2
    left := sortedListToBSTHelper(node, lo, mid - 1)
    newNode := &TreeNode{Val: (*node).Val, Left: left}
    *node = (*node).Next
    newNode.Right = sortedListToBSTHelper(node, mid + 1, hi)
    return newNode
}
