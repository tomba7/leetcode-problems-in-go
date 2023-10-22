/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
    root *TreeNode
    q []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
    cbt := CBTInserter{root: root}
    if root == nil { return cbt }
    cbt.q = append(cbt.q, root)
    for len(cbt.q) != 0 {
        root = cbt.q[0]
        if root.Left == nil {
            return cbt
        } else if root.Right == nil {
            cbt.q = append(cbt.q, root.Left)
            return cbt
        }
        cbt.q = cbt.q[1:]
        cbt.q = append(cbt.q, root.Left, root.Right)
    }
    return cbt
}

func (cbt *CBTInserter) Insert(v int) int {
    newNode := &TreeNode{Val:v}
    if len(cbt.q) == 0 {
        cbt.root = newNode
        cbt.q = append(cbt.q, newNode)
        return -1
    }
    root := cbt.q[0]
    if root.Left == nil {
        root.Left = newNode
    } else {
        root.Right = newNode
        cbt.q = cbt.q[1:]
    }
    cbt.q = append(cbt.q, newNode)
    return root.Val
}

func (cbt *CBTInserter) Get_root() *TreeNode {
    return cbt.root
}

---

type CBTInserter struct {
    root *TreeNode
    q []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
    cbt := CBTInserter{root: root}
    if root == nil { return cbt }
    cbt.q = append(cbt.q, root)
    for len(cbt.q) != 0 {
        root = cbt.q[0]
        if root.Left == nil && root.Right == nil {
            return cbt
        } else if root.Right == nil {
            cbt.q = append(cbt.q, root.Left)
            return cbt
        } else if root.Left == nil {
            panic("Invalid input")
        }
        cbt.q = cbt.q[1:]
        cbt.q = append(cbt.q, root.Left, root.Right)
    }
    return cbt
}

func (cbt *CBTInserter) Insert(v int) int {
    newNode := &TreeNode{Val:v}
    if len(cbt.q) == 0 {
        cbt.root = newNode
        cbt.q = append(cbt.q, newNode)
        return -1
    }
    root := cbt.q[0]
    if root.Left == nil && root.Right == nil {
        root.Left = newNode
    } else if root.Right == nil {
        root.Right = newNode
        cbt.q = cbt.q[1:]
    } else {
        panic("Invalid input")
    }
    cbt.q = append(cbt.q, newNode)
    return root.Val
}

func (cbt *CBTInserter) Get_root() *TreeNode {
    return cbt.root
}
