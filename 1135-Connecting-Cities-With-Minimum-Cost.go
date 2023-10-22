/*
 [Kruskal's Algorithm using Union-Find]
 - Sort the edges by weight in ascending order
 - Traverse the edges and apply union
 - If the vertices are not part of the same group
   perform a union and decrement n. Accumulate the cost
 - If at the end n == 1, we were able to connect all the
   cities, so return the accumulated cost
 - If n != 1 then we were not able to connect all cities
   so return -1
 */
func minimumCost(n int, connections [][]int) int {
    sort.Slice(connections, func(i, j int) bool {
        return connections[i][2] < connections[j][2]
    })
    parent := make([]int, n+1)
    for i := range parent {
        parent[i] = i
    }
    var res int
    for _, conn := range connections {
        p1, p2 := find(conn[0], parent), find(conn[1], parent)
        if p1 != p2 {
            parent[p1] = p2
            res += conn[2]
            n--
        }
    }
    if n != 1 { return -1 }
    return res
}

func find(x int, parent []int) int {
    if parent[x] == x {
        return x
    }
    parent[x] = find(parent[x], parent)
    return parent[x]
}

---

func minimumCost(n int, connections [][]int) int {
    parent := make([]int, n+1)
    for i := 0; i <= n; i++ {
        parent[i] = i
    }
    sort.Slice(connections, func(i, j int) bool {
        return connections[i][2] < connections[j][2]
    })
    var res int
    for _, conn := range connections {
        x, y := conn[0], conn[1]
        if find(x, parent) != find(y, parent) {
            res += conn[2]
            union(x, y, parent, &n)
        }
    }
    if n == 1 {
        return res
    }
    return -1
}

func union(x, y int, parent []int, n *int) {
    px, py := find(x, parent), find(y, parent)
    if px != py {
        parent[px] = py
        *n--
    }
}

func find(x int, parent []int) int {
    if parent[x] == x {
        return parent[x]
    }
    parent[x] = find(parent[x], parent)
    return parent[x]
}
