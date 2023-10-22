func isCousins(root *TreeNode, x int, y int) bool {
    var c cousins
    isCousinsHelper(root, nil, 0, x, y, &c)
    return c.depth1 == c.depth2 && c.parent1 != c.parent2
}

func isCousinsHelper(root, parent *TreeNode, depth, x, y int, c *cousins) {
    if root == nil { return }
    if root.Val == x {
        c.depth1 = depth
        c.parent1 = parent
    } else if root.Val == y {
        c.depth2 = depth
        c.parent2 = parent
    } else {
        isCousinsHelper(root.Left, root, depth+1, x, y, c)
        isCousinsHelper(root.Right, root, depth+1, x, y, c)
    }
}

type cousins struct {
    depth1, depth2 int
    parent1, parent2 *TreeNode
}

---

// Brute force
func isCousins(root *TreeNode, x int, y int) bool {
    d1, p1 := depth(root, nil, x)
    d2, p2 := depth(root, nil, y)
    return d1 == d2 && p1 != p2
}

func depth(root, parent *TreeNode, x int) (int, *TreeNode) {
    if root == nil { return -1, nil }
    if root.Val == x {
        return 0, parent
    }
    left, leftParent := depth(root.Left, root, x)
    if left != -1 {
        return left + 1, leftParent
    }
    right, rightParent := depth(root.Right, root, x)
    if right != -1 {
        return right + 1, rightParent
    }
    return -1, nil
}
