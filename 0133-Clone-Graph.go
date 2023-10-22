func cloneGraph(node *Node) *Node {
    return cloneNode(node, make(map[*Node]*Node))
}

func cloneNode(node *Node, nodeMap map[*Node]*Node) *Node {
    if node == nil { return nil }
    if clone, exist := nodeMap[node]; exist {
        return clone
    }
    clone := &Node{Val: node.Val}
    nodeMap[node] = clone
    for _, neighbor := range node.Neighbors {
        clone.Neighbors = append(clone.Neighbors, cloneNode(neighbor, nodeMap))
    }
    return clone
}

---

func cloneGraph(node *Node) *Node {
    if node == nil { return nil }
    q := list.New()
    cloneMap := make(map[*Node]*Node)
    q.PushBack(node)
    cloneMap[node] = &Node{Val: node.Val}
    for q.Len() != 0 {
        node := q.Front().Value.(*Node)
        q.Remove(q.Front())
        clone := cloneMap[node]
        for _, neighbor := range node.Neighbors {
            clonedNeighbor, exist := cloneMap[neighbor]
            if !exist {
                clonedNeighbor = &Node{Val: neighbor.Val}
                cloneMap[neighbor] = clonedNeighbor
                q.PushBack(neighbor)
            }
            clone.Neighbors = append(clone.Neighbors, clonedNeighbor)
        }
    }
    return cloneMap[node]
}
