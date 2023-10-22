/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var res []int
    for root != nil {
        if root.Left != nil {
            pred := root.Left
            for pred.Right != nil && pred.Right != root {
                pred = pred.Right
            }
            if pred.Right == nil {
                pred.Right = root
                res = append(res, root.Val)
                root = root.Left
            } else {
                pred.Right = nil
                root = root.Right
            }
        } else {
            res = append(res, root.Val)
            root = root.Right
        }
    }
    return res
}

---

func preorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }
    var res []int
    s := []*TreeNode{root}
    for len(s) != 0 {
        root, s = s[len(s)-1], s[:len(s)-1]
        res = append(res, root.Val)
        if root.Right != nil { s = append(s, root.Right) }
        if root.Left != nil { s = append(s, root.Left) }
    }
    return res
}

---

func preorderTraversal(root *TreeNode) []int {
    var res []int
    s := list.New()
    if root != nil {
        s.PushFront(root)
    }
    for s.Len() != 0 {
        node := s.Front().Value.(*TreeNode)
        s.Remove(s.Front())
        res = append(res, node.Val)
        if node.Right != nil { s.PushFront(node.Right) }
        if node.Left != nil { s.PushFront(node.Left) }
    }
    return res
}
