/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    var result []float64
    var q []*TreeNode
    var nodeCount, sum int
    if root != nil {
        q = append(q, root)
        nodeCount++
    }
    numNodesInLevel := nodeCount 
    for len(q) != 0 {
        node := q[0]
        q = q[1:]
        nodeCount--
        sum += node.Val
        if node.Left != nil { q = append(q, node.Left) }
        if node.Right != nil { q = append(q, node.Right) }
        if nodeCount == 0 {
            result = append(result, float64(sum)/float64(numNodesInLevel))
            nodeCount, numNodesInLevel = len(q), len(q)
            sum = 0
        }
    }
    return result
}
