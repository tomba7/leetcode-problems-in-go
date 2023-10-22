func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    var result [][]int
    var levelNodes []int
    q := list.New()
    q.PushBack(root)
    numNodesInLevel := 1
    for q.Len() != 0 {
        node := q.Front().Value.(*TreeNode)
        q.Remove(q.Front())
        levelNodes = append(levelNodes, node.Val)
        numNodesInLevel--
        if node.Left != nil { q.PushBack(node.Left) }
        if node.Right != nil { q.PushBack(node.Right) }
        if numNodesInLevel == 0 {
            numNodesInLevel = q.Len()
            result = append(result, levelNodes)
            levelNodes = nil
        }
    }
    return result
}
