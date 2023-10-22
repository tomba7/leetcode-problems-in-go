/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var res *TreeNode
    lcaHelper(root, p, q, &res)
    return res
}

func lcaHelper(root, p, q *TreeNode, res **TreeNode) int {
    if root == nil { return 0 } 
    left := lcaHelper(root.Left, p, q, res)
    if left == 2 {
        return left
    }
    right := lcaHelper(root.Right, p, q, res)
    if right == 2 {
        return right
    }
    total := left + right
    if root == p || root == q { total++ }
    if total == 2 {
        *res = root
    }
    return total
}
