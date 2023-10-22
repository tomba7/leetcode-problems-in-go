func diameterOfBinaryTree(root *TreeNode) int {
    var res int
    diameterHelper(root, &res)
    return res
}

func diameterHelper(root *TreeNode, res *int) int {
    if root == nil { return 0 }
    left := diameterHelper(root.Left, res)
    right := diameterHelper(root.Right, res)
    *res = max(*res, left + right)
    return max(left, right) + 1
}

func max(x, y int) int { if x > y { return x }; return y }

---

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil { return 0 }
    res, _ := helper(root)
    return res
}

func helper(root *TreeNode) (int, int) {
    if root == nil { return -1, -1 }
    leftDia, leftDepth := helper(root.Left)
    rightDia, rightDepth := helper(root.Right)
    dia := max3(leftDia, rightDia, leftDepth+rightDepth+2)
    depth := max(leftDepth, rightDepth)+1
    return dia, depth
}

func max(x, y int) int { if x > y { return x }; return y }
func max3(x, y, z int) int { return max(max(x, y), z) }

---

func diameterOfBinaryTree(root *TreeNode) int {
    _, dia := diameterHelper(root)
    return dia
}

func diameterHelper(node *TreeNode) (int, int) {
    if node == nil {return 0, 0}
    left, leftDia := diameterHelper(node.Left)
    right, rightDia:= diameterHelper(node.Right)
    return max(left, right) + 1, max(left + right, max(leftDia, rightDia))
}

func max(x, y int) int {if x > y {return x}; return y}

---

func diameterOfBinaryTree(root *TreeNode) int {
    var dia int
    diameterHelper(root, &dia)
    return dia
}

func diameterHelper(root *TreeNode, dia *int) int {
    if root == nil { return 0 }
    left := diameterHelper(root.Left, dia)
    right := diameterHelper(root.Right, dia)
    *dia = max(*dia, left + right)
    return max(left, right) + 1
}

func max(x, y int) int { if x > y {return x}; return y }

---

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil { return 0 }
    _, res := diameterHelper(root)
    return res
}

func diameterHelper(root *TreeNode) (int, int) {
    if root == nil { return -1, -1 }
    ltDepth, ltDia := diameterHelper(root.Left)
    rtDepth, rtDia := diameterHelper(root.Right)
    return max(ltDepth, rtDepth)+1, max(max(ltDia, rtDia), ltDepth+rtDepth+2)
}

func max(x, y int) int { if x > y {return x}; return y }
