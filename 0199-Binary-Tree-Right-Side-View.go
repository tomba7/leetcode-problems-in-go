func rightSideView(root *TreeNode) []int {
    var res []int
    helper(root, 0, &res)
    return res
}

func helper(root *TreeNode, level int, res *[]int) {
    if root == nil { return }
    if level == len(*res) {
        *res = append(*res, root.Val)
    }
    helper(root.Right, level+1, res)
    helper(root.Left, level+1, res)
}

---

func rightSideView(root *TreeNode) []int {
    if root == nil { return nil }
    q := list.New()
    q.PushBack(root)
    levelSize := 1
    var res []int
    for q.Len() != 0 {
        curr := q.Remove(q.Front()).(*TreeNode)
        if curr.Left != nil { q.PushBack(curr.Left) }
        if curr.Right != nil { q.PushBack(curr.Right) }
        levelSize--
        if levelSize == 0 {
            res = append(res, curr.Val)
            levelSize = q.Len()
        }
    }
    return res
}

---

func rightSideView(root *TreeNode) []int {
    if root == nil { return nil }
    result, queue, numNodesAtLevel := []int{}, []*TreeNode{}, 1
    queue = append(queue, root)
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        numNodesAtLevel--
        if node.Left != nil { queue = append(queue, node.Left) }
        if node.Right != nil { queue = append(queue, node.Right) }
        if numNodesAtLevel == 0 {
            result = append(result, node.Val)
            numNodesAtLevel = len(queue)
        }
    }
    return result
}
