func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    q := list.New()
    q.PushBack(beginWord)
    dict := newDictionary(wordList)
    seen := map[string]bool{beginWord: true}
    distance := map[string]int{beginWord: 1}
    for q.Len() != 0 {
        curr := q.Remove(q.Front()).(string)
        if curr == endWord { break }
        for _, next := range neighbors(curr) {
            if !dict[next] || seen[next] { continue }
            seen[next] = true
            distance[next] = distance[curr] + 1
            q.PushBack(next)
        }
    }
    var res [][]string
    dfs(beginWord, endWord, dict, distance, nil, &res)
    return res
}

func dfs(curr, endWord string, dict map[string]bool,
         distance map[string]int, path []string, res *[][]string) {
    path = append(path, curr)
    if curr == endWord {
        *res = append(*res, append([]string{}, path...))
        return
    }
    for _, next := range neighbors(curr) {
        if distance[next] != distance[curr] + 1 {
            continue
        }
        dfs(next, endWord, dict, distance, path, res)
    }
}

func neighbors(word string) []string {
    var res []string
    m := len(word)
    buf := []byte(word)
    for i := 0; i < m; i++ {
        curr := buf[i]
        for ch := byte('a'); ch <= 'z'; ch++ {
            if curr == ch { continue }
            buf[i] = ch
            res = append(res, string(buf))
        }
        buf[i] = curr
    }
    return res
}

func newDictionary(wordList []string) map[string]bool {
    dict := map[string]bool{}
    for _, word := range wordList {
        dict[word] = true
    }
    return dict
}
