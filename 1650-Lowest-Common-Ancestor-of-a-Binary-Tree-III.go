/*
 - Starting at both p and q, traverse until Parent != nil
 - At the end we'll have length m and n of both paths.
 - Subtract both the path length to get x
 - Traverse the longer path by x.
 - From here traverse both pointers in tandem.
 - When the parents become the same, that is the LCA node
 */
func lowestCommonAncestor(p *Node, q *Node) *Node {
    m, n := pathLength(p), pathLength(q)
    if n > m { return lowestCommonAncestor(q, p) }
    for i := 0; i < m-n; i++ {
        p = p.Parent
    }
    for p != nil && q != nil {
        if p == q { return p }
        p, q = p.Parent, q.Parent
    }
    return nil
}

func pathLength(curr *Node) int {
    var i int
    for curr != nil {
        curr = curr.Parent
        i++
    }
    return i
}
