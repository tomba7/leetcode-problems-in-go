/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil { return nil }
    result := [][]int{}
    q, levelNodes := []*TreeNode{}, []int{}
    q = append(q, root)
    numNodesInLevel := len(q)
    for len(q) != 0 {
        node := q[0]
        q = q[1:]
        numNodesInLevel--
        if node.Left != nil { q = append(q, node.Left) }
        if node.Right != nil { q = append(q, node.Right) }
        levelNodes = append(levelNodes, node.Val)
        if numNodesInLevel == 0 {
            numNodesInLevel = len(q)
            result = append(result, levelNodes)
            levelNodes = nil
        }
    }
    return result
}
