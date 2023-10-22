/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    q := list.New()
    q.PushBack(root)
    terminating := false
    for q.Len() != 0 {
        curr := q.Front().Value.(*TreeNode)
        q.Remove(q.Front())
        if curr.Left == nil && curr.Right != nil {
            return false
        } else if terminating && curr.Left != nil {
            return false
        } else if curr.Right == nil {
            terminating = true
        }
        if curr.Left != nil {
            q.PushBack(curr.Left)
        }
        if curr.Right != nil {
            q.PushBack(curr.Right)
        }
    }
    return true
}

---

func isCompleteTree(root *TreeNode) bool {
    if root == nil { return false }
    q := []*TreeNode{root}
    var endMarker bool
    for len(q) != 0 {
        root, q = q[0], q[1:]
        if endMarker && (root.Left != nil || root.Right != nil) {
            return false
        } else if root.Right == nil {
            endMarker = true
        } else if root.Left == nil {
            return false
        }
        if root.Left != nil { q = append(q, root.Left) }
        if root.Right != nil { q = append(q, root.Right) }
    }
    return true
}
