func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    var res []*TreeNode
    postOrder(root, make(map[string]int), &res)
    return res
}

func postOrder(root *TreeNode, table map[string]int, res *[]*TreeNode) string {
    if root == nil { return "#" }
    left := postOrder(root.Left, table, res)
    right := postOrder(root.Right, table, res)
    serial := fmt.Sprintf("%d,%s,%s", root.Val, left, right)
    table[serial]++
    if table[serial] == 2 {
        *res = append(*res, root)
    }
    return serial
}

---

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    var res []*TreeNode
    id := 1
    serialToId := map[string]int{}
    idToCount := map[int]int{}
    postOrder(root, serialToId, idToCount, &id, &res)
    return res
}

func postOrder(root *TreeNode, serialToId map[string]int,
               idToCount map[int]int, id *int, res *[]*TreeNode) int {
    if root == nil { return 0 }
    left := postOrder(root.Left, serialToId, idToCount, id, res)
    right := postOrder(root.Right, serialToId, idToCount, id, res)
    serial := fmt.Sprintf("%d,%d,%d", root.Val, left, right)
    serialId, exist := serialToId[serial]
    if !exist {
        serialId = *id
        serialToId[serial] = *id
        *id++
    }
    idToCount[serialId]++
    if idToCount[serialId] == 2 {
        *res = append(*res, root)
    }
    return serialId
}

---

func findDuplicateSubtrees(root *TreeNode) (result []*TreeNode) {
    m := map[string][]*TreeNode{}
    traverse(root, m)
    for _, v := range m {
        if len(v) > 1 { result = append(result, v[0]) }
    }
    return result
}

func traverse(root *TreeNode, m map[string][]*TreeNode) (s string) {
    if root == nil { return "x" }
    s = strconv.Itoa(root.Val) + traverse(root.Left, m) + traverse(root.Right, m)
    m[s] = append(m[s], root)
    return s
}
