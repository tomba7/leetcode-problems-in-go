func alienOrder(words []string) string {
    var res []byte
    graph := newGraph(words)
    seen := make(map[byte]int)
    for ch := range graph {
        if !dfs(ch, graph, seen, &res) { return "" }
    }
    reverse(res)
    return string(res)
}

func dfs(ch byte, graph map[byte][]byte, seen map[byte]int, res *[]byte) bool {
    if seen[ch] == 2 {
        return true
    } else if seen[ch] == 1 {
        // Cycle detected
        return false
    }
    seen[ch] = 1    // Visiting
    for _, neighbor := range graph[ch] {
        if !dfs(neighbor, graph, seen, res) { return false }
    }
    seen[ch] = 2    // Done
    *res = append(*res, ch)
    return true
}

func newGraph(words []string) map[byte][]byte {
    graph := make(map[byte][]byte)
    for _, word := range words {
        for i := 0; i < len(word); i++ {
            if _, exist := graph[word[i]]; exist { continue }
            graph[word[i]] = []byte{}
        }
    }
    for i := 0; i < len(words)-1; i++ {
        word1, word2 := words[i], words[i+1]
        var j, k int
        for ; j < len(word1) && k < len(word2); j, k = j+1, k+1 {
            if word1[j] == word2[j] { continue }
            graph[word1[j]] = append(graph[word1[j]], word2[j])
            break
        }
        if j < len(word1) && k == len(word2) { return nil }
    }
    return graph
}

func reverse(s []byte) {
    for i, j := 0, len(s)-1; i < j; {
        s[i], s[j] = s[j], s[i]
        i++
        j--
    }
}

---

// 2:42 PM
/*
- create graph
- DFS with coloring (cycle detection)
 */
func alienOrder(words []string) string {
    graph := newGraph(words)
    seen := make(map[byte]int)
    var res []byte
    for ch := range graph {
        if !dfs(graph, seen, ch, &res) {
            return ""
        }
    }
    return reverse(res)
}

// newGraph
func newGraph(words []string) map[byte][]byte {
    graph := make(map[byte][]byte)
    for _, word := range words {
        for i := 0; i < len(word); i++ {
            if _, exist := graph[word[i]]; exist { continue }
            graph[word[i]] = []byte{}
        }
    }
    for i := 0; i < len(words)-1; i++ {
        word1, word2 := words[i], words[i+1]
        var j int
        for ; j < len(word1) && j < len(word2); j++ {
            if word1[j] == word2[j] { continue }
            graph[word1[j]] = append(graph[word1[j]], word2[j])
            break
        }
        if j < len(word1) && j == len(word2) { return nil }
    }
    return graph
}

// func dfs
func dfs(graph map[byte][]byte, seen map[byte]int, ch byte, res *[]byte) bool {
    if seen[ch] == 2 {
        return true
    } else if seen[ch] == 1 {
        return false
    }
    seen[ch] = 1
    for _, next := range graph[ch] {
        if !dfs(graph, seen, next, res) { return false }
    }
    seen[ch] = 2
    *res = append(*res, ch)
    return true
}

// reverse()
func reverse(s []byte) string {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return string(s)
}
