/*
     1  2  3
 1 [-1, 1, 0]
 2 [ 1,-1, 0]
 3 [ 0, 0, 1]
 
 - Traverse all cities from 1 to n
 - Perform DFS on each on of them and increment the
   number of provinces once they return
 - If a city as been visited already bypass the DFS
 */

func findCircleNum(isConnected [][]int) int {
    var res int
    n := len(isConnected)
    seen := make(map[int]bool)
    for i := 0; i < n; i++ {
        if seen[i] { continue }
        dfs(isConnected, i, seen)
        res++
    }
    return res
}

func dfs(graph [][]int, i int, seen map[int]bool) {
    n := len(graph)
    seen[i] = true
    for j := 0; j < n; j++ {
        if j == i { continue }
        if graph[i][j] == 1 && !seen[j] {
            dfs(graph, j, seen)
        }
    }
}

---

func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    seen := make([]bool, n)
    var res int
    q := list.New()
    for i := 0; i < n; i++ {
        if seen[i] { continue }
        q.PushBack(i)
        for q.Len() != 0 {
            j := q.Remove(q.Front()).(int)
            seen[j] = true
            for k := 0; k < n; k++ {
                if j == k || seen[k] || isConnected[j][k] == 0 {
                    continue
                }
                q.PushBack(k)
            }
        }
        res++
    }
    return res
}

---

func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    seen := make([]bool, n)
    var res int
    for city := 0; city < n; city++ {
        if seen[city] {
            continue
        }
        dfs(city, isConnected, seen)
        res++
    }
    return res
}

func dfs(city int, isConnected [][]int, seen []bool) {
    n := len(isConnected)
    seen[city] = true
    for nextCity := 0; nextCity < n; nextCity++ {
        if nextCity == city || isConnected[city][nextCity] == 0 || seen[nextCity] {
            continue
        }
        dfs(nextCity, isConnected, seen)
    }
}
