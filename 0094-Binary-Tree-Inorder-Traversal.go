/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var res []int
    for root != nil {
        if root.Left != nil {
            pred := root.Left
            for pred.Right != nil && pred.Right != root {
                pred = pred.Right
            }
            if pred.Right == nil {
                pred.Right = root
                root = root.Left
            } else {
                pred.Right = nil
                res = append(res, root.Val)
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

func inorderTraversal(root *TreeNode) []int {
    var res []int
    curr := root
    s := list.New()
    for curr != nil || s.Len() != 0 {
        for curr != nil {
            s.PushBack(curr)
            curr = curr.Left
        }
        curr = s.Back().Value.(*TreeNode)
        s.Remove(s.Back())
        res = append(res, curr.Val)
        curr = curr.Right
    }
    return res
}

---

func inorderTraversal(root *TreeNode) []int {
    var res []int
    curr := root
    s := list.New()
    for ; curr != nil; curr = curr.Left {
       s.PushBack(curr)
    }
    for s.Len() != 0 {
        curr = s.Back().Value.(*TreeNode)
        s.Remove(s.Back())
        res = append(res, curr.Val)
        curr = curr.Right
        for ; curr != nil; curr = curr.Left {
            s.PushBack(curr)
        }
    }
    return res
}

---

func inorderTraversal(root *TreeNode) []int {
    var res []int
    var s []*TreeNode
    for len(s) != 0 || root != nil {
        for root != nil {
            s = append(s, root)
            root = root.Left
        }
        root, s = s[len(s)-1], s[:len(s)-1]
        res = append(res, root.Val)
        root = root.Right
    }
    return res
}
