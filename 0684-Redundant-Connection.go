// DFS Approach
func findRedundantConnection(edges [][]int) []int {
    graph := map[int][]int{}
    for _, edge := range edges {
        seen := map[int]bool{}
        src, dst := edge[0], edge[1]
        if graph[src] != nil && graph[dst] != nil && dfs(src, dst, graph, seen) {
            return edge
        }
        graph[src] = append(graph[src], dst)
        graph[dst] = append(graph[dst], src)
    }
    return nil
}

func dfs(src, dst int, graph map[int][]int, seen map[int]bool) bool {
    if seen[src] { return false }
    seen[src] = true
    if src == dst { return true }
    for _, nei := range graph[src] {
        if dfs(nei, dst, graph, seen) { return true }
    }
    return false
}

---

func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    parent := make([]int, n+1)
    for i := 1; i <= n; i++ {
        parent[i] = i
    }
    for _, edge := range edges {
        p1, p2 := find(parent, edge[0]), find(parent, edge[1])
        if p1 != p2 {
            parent[p1] = p2
        } else {
            return edge
        }
    }
    return nil
}

func find(parent []int, x int) int {
    for parent[x] != x {
        parent[x] = parent[parent[x]]
        x = parent[x]
    }
    return x
}

---

// Union Find Approach
func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    parent := make([]int, n+1)
    for i := 0; i < len(parent); i++ {
        parent[i] = i
    }
    rank := make([]int, n+1)
    for i := 0; i < len(rank); i++ {
        rank[i] = 1
    }
    for _, edge := range edges {
        if !union(edge[0], edge[1], parent, rank) {
            return edge
        }
    }
    return nil
}

func find(x int, parent []int) int {
    p := parent[x]
    for p != parent[p] {
        // The following 'called Path Compression' can be used as well
        // parent[p] = parent[parent[p]]
        p = parent[p]
    }
    return p
}

func union(x1, x2 int, parent, rank []int) bool {
    p1, p2 := find(x1, parent), find(x2, parent)
    if p1 == p2 { return false }
    if rank[p1] > rank[p2] {
        parent[p2] = p1
        rank[p1] += rank[p2]
    } else {
        parent[p1] = p2
        rank[p2] += rank[p1]
    }
    return true
}

---

func findRedundantConnection(edges [][]int) []int {
    n := len(edges) + 1
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    rank := make([]int, n)
    for _, edge := range edges {
        if !union(edge[0], edge[1], parent, rank) { return edge }
    }
    return nil
}

func union(x, y int, parent, rank []int) bool {
    rootX, rootY := find(x, parent), find(y, parent)
    if rootX == rootY {
        return false
    } else if rank[rootX] < rank[rootY] {
        parent[rootX] = rootY
    } else if rank[rootX] > rank[rootY] {
        parent[rootY] = rootX
    } else {
        parent[rootY] = rootX
        rank[rootX]++
    }
    return true
}

func find(x int, parent []int) int {
    if parent[x] != x {
        parent[x] = find(parent[x], parent)
    }
    return parent[x]
}
