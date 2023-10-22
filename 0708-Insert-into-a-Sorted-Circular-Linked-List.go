func insert(node *Node, x int) *Node {
    if node == nil {
        node := &Node{Val: x}
        node.Next = node
        return node
    }
    curr, next := node, node.Next
    for curr.Next != node {
        if x >= curr.Val && x <= next.Val ||
            curr.Val > next.Val && (x < next.Val || x > curr.Val) {
            curr.Next = &Node{Val: x, Next: next}
            return node
        }
        curr, next = curr.Next, next.Next
    }
    curr.Next = &Node{Val: x, Next: next}
    return node
}

---

func insert(node *Node, x int) *Node {
    if node == nil {
        newNode := &Node{Val: x}
        newNode.Next = newNode
        return newNode
    }
    curr, next := node, node.Next
    for next != node {
        if x >= curr.Val && x <= next.Val ||
           curr.Val > next.Val && (x > curr.Val || x < next.Val) {
            break
        }
        curr = curr.Next
        next = next.Next
    }
    curr.Next = &Node{
        Val: x, Next: next,
    }
    return node
}

---

func insert(node *Node, x int) *Node {
    if node == nil {
        node := &Node{Val: x}
        node.Next = node
        return node
    }
    var inserted bool
    curr, next := node, node.Next
    for {
        if curr.Val <= x && x <= next.Val ||
            curr.Val > next.Val && (x < next.Val || x > curr.Val) {
            curr.Next = &Node{Val: x, Next: next}
            inserted = true
            return node
        }
        curr, next = curr.Next, next.Next
        if curr == node { break }
    }
    if !inserted {
        curr.Next = &Node{Val: x, Next: next}
    }
    return node
}
