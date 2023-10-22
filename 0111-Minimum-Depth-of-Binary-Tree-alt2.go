/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil { return 0 }
    q := list.New()
    q.PushBack(root)
    numNodesInLevel := 1
    depth := 1
    for q.Len() != 0 {
        node := q.Front().Value.(*TreeNode)
        q.Remove(q.Front())
        if node.Left == nil && node.Right == nil { return depth }
        if node.Left != nil { q.PushBack(node.Left) }
        if node.Right != nil { q.PushBack(node.Right) }
        numNodesInLevel--
        if numNodesInLevel == 0 {
            depth++
            numNodesInLevel = q.Len()
        }
    }
    return depth
}
