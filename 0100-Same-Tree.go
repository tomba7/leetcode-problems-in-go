/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil { return true }
    if p == nil || q == nil || p.Val != q.Val { return false }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

---

func isSameTree(p *TreeNode, q *TreeNode) bool {
    pserial, qserial := serialize(p), serialize(q)
    return pserial == qserial
}

func serialize(root *TreeNode) string {
    var res []byte
    serializeHelper(root, &res)
    return string(res)
}

func serializeHelper(root *TreeNode, res *[]byte) {
    if root == nil {
        *res = append(*res, ',', '#')
        return
    }
    *res = append(*res, fmt.Sprintf("%d,", root.Val)...)
    serializeHelper(root.Left, res)
    serializeHelper(root.Right, res)
}
