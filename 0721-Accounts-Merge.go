/*
    DFS Solution
    - Build a graph. For each account,
    -   strip the name
    -   make an edge from the first email to rest of the emails
    -     if there is only one email, dont put any edges
    -   Store a mapping from email to names
    -
    - Go through the graph and perform DFS on each. Skip an email if its seen already
    - After the DFS, append an empty string and sort the result.
    - Put the name on the first spot where the empty string is
    - Append to the final result and return
*/
func accountsMerge(accounts [][]string) [][]string {
    emailToNames := make(map[string]string)
    graph := make(map[string][]string)
    for _, account := range accounts {
        name := account[0]
        if _, contains := graph[account[1]]; !contains {
            graph[account[1]] = []string{}
            emailToNames[account[1]] = name
        }
        for i := 2; i < len(account); i++ {
            graph[account[1]] = append(graph[account[1]], account[i])
            graph[account[i]] = append(graph[account[i]], account[1])
            emailToNames[account[i]] = name
        }
    }
    seen := make(map[string]bool)
    var res [][]string
    for email := range graph {
        if seen[email] { continue }
        path := []string{emailToNames[email]}
        dfs(email, graph, seen, &path)
        sort.Strings(path[1:])
        res = append(res, path)
    }
    return res
}

func dfs(email string, graph map[string][]string, seen map[string]bool, path *[]string) {
    seen[email] = true
    *path = append(*path, email)
    for _, next := range graph[email] {
        if seen[next] { continue }
        dfs(next, graph, seen, path)
    }
}

---

func accountsMerge(accounts [][]string) [][]string {
    graph := make(map[string][]string)
    emailToName := make(map[string]string)
    for _, acc := range accounts {          // O(n)
        name := acc[0]
        for i := 1; i < len(acc); i++ {     // O(k)
            emailToName[acc[i]] = name
            graph[acc[1]] = append(graph[acc[1]], acc[i])
            if i == 1 { continue }
            graph[acc[i]] = append(graph[acc[i]], acc[1])
        }
    }                                       // O(nk)
    visited := make(map[string]bool)
    var res [][]string
    for email := range graph {              // O(n) worst case
        if visited[email] { continue }
        list := []string{""}
        dfs(email, graph, visited, &list)   // O(k) worse case
        sort.Strings(list)                  // O(k log k)
        list[0] = emailToName[email]
        res = append(res, list)
    }
    return res                              // O(nk) + O(nk logk)
}

func dfs(email string, graph map[string][]string, visited map[string]bool, list *[]string) {
    visited[email] = true
    *list = append(*list, email)
    for _, neighbor := range graph[email] {
        if visited[neighbor] { continue }
        dfs(neighbor, graph, visited, list)
    }
}
