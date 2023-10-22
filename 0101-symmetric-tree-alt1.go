/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    q := list.New()
    q.PushBack(root)
    q.PushBack(root)
    for q.Len() != 0 {
        left := q.Front().Value.(*TreeNode)
        q.Remove(q.Front())
        right := q.Front().Value.(*TreeNode)
        q.Remove(q.Front())
        if left == nil && right == nil {
            continue
        } else if left == nil || right == nil || left.Val != right.Val {
            return false
        }
        q.PushBack(left.Left)
        q.PushBack(right.Right)
        q.PushBack(left.Right)
        q.PushBack(right.Left)
    }
    return true
}
