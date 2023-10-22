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

    q := []*TreeNode{}
    q = append(q, root)
    rightMost := root
    depth := 1

    for len(q) != 0 {
        node := q[0]
        q = q[1:]
        if node.Left == nil && node.Right == nil { return depth }
        if node.Left != nil { q = append(q, node.Left) }
        if node.Right != nil { q = append(q, node.Right) }

        if node == rightMost {
            depth++
            if node.Right != nil {
                rightMost = node.Right
            } else {
                rightMost = node.Left
            }
        }
    }
    return depth
}

func min(x, y int) int { if x < y { return x }; return y }
