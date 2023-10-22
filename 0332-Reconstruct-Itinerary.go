/*
 9:51 AM
 - Create a graph[src] = [dst1, dst2, dst3] where dst's are lex sorted
 - Create a seen[src] = [false, false, false]
 - Backtracking/DFS
 - if i == len(tickets) + 1, we have used all tickets succesfully
 */

func findItinerary(tickets [][]string) []string {
    n := len(tickets)
    graph := newGraph(tickets)
    seen := map[string][]bool{}
    for src, dstList := range graph {
        seen[src] = make([]bool, len(dstList))
    }
    var res []string
    path := make([]string, n+1)
    findItineraryHelper(graph, seen, "JFK", 0, path, &res)
    return res
}

func findItineraryHelper(graph map[string][]string, seen map[string][]bool,
                         src string, i int, path []string, res *[]string) bool {
    path[i] = src
    if i == len(path)-1 {
        *res = path
        return true
    }
    for j, dst := range graph[src] {
        if seen[src][j] { continue }
        seen[src][j] = true
        if findItineraryHelper(graph, seen, dst, i+1, path, res) {
            return true
        }
        seen[src][j] = false
    }
    return false
}

func newGraph(tickets [][]string) map[string][]string {
    graph := map[string][]string{}
    for _, ticket := range tickets {
        src, dst := ticket[0], ticket[1]
        graph[src] = append(graph[src], dst)
        if _, exist := graph[dst]; !exist { graph[dst] = []string{} }
    }
    for _, dstList := range graph {
        sort.Strings(dstList)
    }
    return graph
}

---

func findItinerary(tickets [][]string) []string {
    graph := newGraph(tickets)
    seen := make(map[string][]bool)
    for from, destinations := range graph {
        seen[from] = make([]bool, len(destinations))
    }
    _, res := backtrack("JFK", make([]string, 0), graph, seen, len(tickets))
    return res
}

func backtrack(curr string, itinerary []string, graph map[string][]string,
               seen map[string][]bool, n int) (bool, []string) {
    itinerary = append(itinerary, curr)
    if len(itinerary) == n + 1 {
        return true, itinerary
    }
    for i := 0; i < len(graph[curr]); i++ {
        next := graph[curr][i]
        if seen[curr][i] { continue }
        seen[curr][i] = true
        found, res := backtrack(next, itinerary, graph, seen, n)
        seen[curr][i] = false
        if found { return true, res }
    }
    return false, nil
}

func newGraph(tickets [][]string) map[string][]string {
    graph := make(map[string][]string)
    for _, ticket := range tickets {
        from, to := ticket[0], ticket[1]
        graph[from] = append(graph[from], to)
    }
    for _, destinations := range graph {
        sort.Strings(destinations)
    }
    return graph
}
