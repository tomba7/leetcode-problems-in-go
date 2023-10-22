func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    parentTable := make(map[*TreeNode]*TreeNode)
    findParents(root, nil, parentTable)
    q := list.New()
    q.PushBack(&Element{node: target, distance: 0})
    seen := make(map[*TreeNode]bool)
    seen[target] = true
    var res []int
    for q.Len() != 0 {
        element := q.Remove(q.Front()).(*Element)
        node, distance := element.node, element.distance
        if distance == k {
            res = append(res, node.Val)
        }
        if node.Left != nil && !seen[node.Left] {
            seen[node.Left] = true
            q.PushBack(&Element{node: node.Left, distance: distance + 1})
        }
        if node.Right != nil && !seen[node.Right] {
            seen[node.Right] = true
            q.PushBack(&Element{node: node.Right, distance: distance + 1})
        }
        parent := parentTable[node]
        if parent != nil && !seen[parent] {
            seen[parent] = true
            q.PushBack(&Element{node: parent, distance: distance + 1})
        }
    }
    return res
}

func findParents(root, parent *TreeNode, parentTable map[*TreeNode]*TreeNode) {
    if root == nil { return }
    parentTable[root] = parent
    findParents(root.Left, root, parentTable)
    findParents(root.Right, root, parentTable)
}

type Element struct {
    node *TreeNode
    distance int
}

---

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    parentIndex := newParentIndex(root)
    seen := make([]bool, 501)
    q := list.New()
    q.PushBack(newPair(target, 0))
    var res []int
    for q.Len() != 0 {
        pair := q.Remove(q.Front()).(*Pair)
        node, dist := pair.node, pair.dist
        seen[node.Val] = true
        if dist == k {
            res = append(res, node.Val)
        }
        if node.Left != nil && !seen[node.Left.Val] {
            q.PushBack(newPair(node.Left, dist + 1))
        }
        if node.Right != nil && !seen[node.Right.Val] {
            q.PushBack(newPair(node.Right, dist + 1))
        }
        parent := parentIndex[node.Val]
        if parent != nil && !seen[parent.Val] {
            q.PushBack(newPair(parent, dist + 1))
        }
    }
    return res
}

type Pair struct {
    node *TreeNode
    dist int
}

func newPair(node *TreeNode, dist int) *Pair {
    return &Pair{node: node, dist: dist}
}

func newParentIndex(root *TreeNode) []*TreeNode {
    parentIndex := make([]*TreeNode, 501)
    newParentIndexHelper(root, nil, parentIndex)
    return parentIndex
}

func newParentIndexHelper(root, parent *TreeNode, parentIndex []*TreeNode) {
    if root == nil { return }
    parentIndex[root.Val] = parent
    newParentIndexHelper(root.Left, root, parentIndex)
    newParentIndexHelper(root.Right, root, parentIndex)
}
