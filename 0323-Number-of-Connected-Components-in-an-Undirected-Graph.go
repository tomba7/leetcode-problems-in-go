// DFS
func countComponents(n int, edges [][]int) int {
    var count int
    graph := newGraph(n, edges)
    seen := make([]bool, n)
    for i := 0; i < n; i++ {
        if seen[i] { continue }
        dfs(i, graph, seen)
        count++
    }
    return count
}

func dfs(i int, graph [][]int, seen []bool) {
    seen[i] = true
    for _, neighbor := range graph[i] {
        if seen[neighbor] { continue }
        dfs(neighbor, graph, seen)
    }
}

func newGraph(n int, edges [][]int) [][]int {
    graph := make([][]int, n)
    for _, edge := range edges {
        src, dst := edge[0], edge[1]
        graph[src] = append(graph[src], dst)
        graph[dst] = append(graph[dst], src)
    }
    return graph
}

---

// Union-Find
func countComponents(n int, edges [][]int) int {
    parent := make([]int, n)
    for i := 0; i < n; i++ { parent[i] = i }
    for _, edge := range edges {
        p1 := find(parent, edge[0])
        p2 := find(parent, edge[1])
        if p1 != p2 {
            parent[p1] = p2
            n--
        }
    }
    return n
}

func find(parent []int, x int) int {
    for parent[x] != x {
        parent[x] = parent[parent[x]]
        x = parent[x]
    }
    return x
}

---

func countComponents(n int, edges [][]int) int {
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    for _, edge := range edges {
        p1, p2 := find(edge[0], parent), find(edge[1], parent)
        if p1 != p2 {
            parent[p1] = p2
            n--
        }
    }
    return n
}

func find(x int, parent []int) int {
    if x == parent[x] { return x }
    parent[x] = find(parent[x], parent)
    return parent[x]
}
