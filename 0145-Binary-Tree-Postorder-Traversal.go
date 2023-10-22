/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var res []int
    var s []*TreeNode
    curr := root
    for curr != nil || len(s) != 0 {
        for curr != nil {
            if curr.Right != nil {
                s = append(s, curr.Right)
            }
            s = append(s, curr)
            curr = curr.Left
        }
        curr, s = s[len(s)-1], s[:len(s)-1]
        if curr.Right == nil {
            res = append(res, curr.Val)
            curr = nil
        } else {
            if len(s) != 0 && curr.Right == s[len(s)-1] {
                // Pop off the top of the stack and push curr
                s[len(s)-1] = curr
                curr = curr.Right
            } else {
                res = append(res, curr.Val)
                curr = nil
            }
        }
    }
    return res
}

---

func postorderTraversal(root *TreeNode) []int {
    var res []int
    s := list.New()
    for curr := root; curr != nil || s.Len() != 0; {
        for curr != nil {
            if curr.Right != nil {
                s.PushBack(curr.Right)
            }
            s.PushBack(curr)
            curr = curr.Left
        }
        curr = s.Remove(s.Back()).(*TreeNode)
        if curr.Right == nil {
            res = append(res, curr.Val)
            curr = nil
        } else {
            if s.Len() != 0 && s.Back().Value.(*TreeNode) == curr.Right {
                s.Remove(s.Back())
                s.PushBack(curr)
                curr = curr.Right
            } else {
                res = append(res, curr.Val)
                curr = nil
            }
        }
    }
    return res
}

---

func postorderTraversal(root *TreeNode) []int {
    var res []int
    if root == nil { return res }
    s := []*TreeNode{root}
    for len(s) != 0 {
        root = s[len(s)-1]
        if root.Left == nil && root.Right == nil {
            res = append(res, root.Val)
            s = s[:len(s)-1]
        } else {
            if root.Right != nil {
                s = append(s, root.Right)
                root.Right = nil
            }
            if root.Left != nil {
                s = append(s, root.Left)
                root.Left = nil
            }
        }
    }
    return res
}

---
// Avoid this approach

/*
 * https://techgeekyan.blogspot.com/2017/08/leetcode-145-binary-tree-postorder.html
 */
func postorderTraversal(root *TreeNode) []int {
    var result []int
    dummy := &TreeNode{Left: root}
    curr := dummy
    for curr != nil {
        if curr.Left == nil {
            curr = curr.Right
            continue
        }
        pred := curr.Left
        for pred.Right != nil && pred.Right != curr {
            pred = pred.Right
        }
        if pred.Right == nil {
            pred.Right = curr
            curr = curr.Left
        } else {
            reverse(curr.Left, pred)
            for tmp := pred; ; tmp = tmp.Right {
                result = append(result, tmp.Val)
                if tmp == curr.Left { break }
            }
            reverse(pred, curr.Left)
            pred.Right = nil
            curr = curr.Right
        }
    }
    return result
}

func reverse(from, to *TreeNode) {
    x, y := from, from.Right
    for x != to {
        x, y, y.Right = y, y.Right, x
    }
}
