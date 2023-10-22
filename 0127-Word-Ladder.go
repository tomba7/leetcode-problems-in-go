func ladderLength(beginWord, endWord string, wordList []string) int {
    q := list.New()
    q.PushBack(Pair{word: beginWord, length: 1})
    visited := map[string]bool{}
    dict := newDictionary(wordList)
    for q.Len() != 0 {
        pair := q.Front().Value.(Pair)
        q.Remove(q.Front())
        for _, neighbor := range getNeighbors(pair.word) {
            if !dict[neighbor] || visited[neighbor] { continue }
            if neighbor == endWord { return pair.length+1 }
            visited[neighbor] = true
            q.PushBack(Pair{word: neighbor, length: pair.length+1})
        }
    }
    return 0
}

type Pair struct {
    word string
    length int
}

func newDictionary(words []string) map[string]bool {
    dict := make(map[string]bool)
    for _, word := range words {
        dict[word] = true
    }
    return dict
}

func getNeighbors(word string) []string {
    bslice := []byte(word)
    var neighbors []string
    for i := 0; i < len(word); i++ {
        orig := bslice[i]
        for c := byte('a'); c <= 'z'; c++ {
            if c == orig { continue }
            bslice[i] = c
            neighbors = append(neighbors, string(bslice))
        }
        bslice[i] = orig
    }
    return neighbors
}

---

/* 10:30 PM, 10:55 PM
- Create a function to generate all possible neighbors given a word of length m characters
  These neighbors will differ from the current word by just a single character
- Create a set from wordList for O(1) lookup
- Create a seen table
- Perform DFS starting at beginWord
- For every element in the queue, get all possible neighbors and filter them
  based on the set
 */

func ladderLength(beginWord string, endWord string, wordList []string) int {
    q := list.New()
    q.PushBack(beginWord)
    dict := newDictionary(wordList)
    seen := map[string]bool{beginWord: true}
    dist := map[string]int{beginWord: 1}
    for q.Len() != 0 {
        curr := q.Remove(q.Front()).(string)
        if curr == endWord {
            return dist[endWord]
        }
        for _, next := range neighbors(curr) {
            if !dict[next] || seen[next] { continue }
            seen[next] = true
            dist[next] = dist[curr] + 1
            q.PushBack(next)
        }
    }
    return 0
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

---

func ladderLength(beginWord string, endWord string, wordList []string) int {
	graph := newGraph(wordList)
	q := list.New()
	q.PushBack(pair{word: beginWord, length: 1})
	visited := map[string]bool{beginWord: true}
	for q.Len() != 0 {
		p := q.Front().Value.(pair)
		q.Remove(q.Front())
		for i := 0; i < len(p.word); i++ {
			intermediate := p.word[:i] + "*" + p.word[i+1:]
			for _, word := range graph[intermediate] {
				if word == endWord {
					return p.length + 1
				}
				if !visited[word] {
					visited[word] = true
					q.PushBack(pair{word: word, length: p.length+1})
				}
			}
		}
	}
	return 0
}

func newGraph(words []string) map[string][]string {
	graph := make(map[string][]string)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			curr := word[:i] + "*" + word[i+1:]
			graph[curr] = append(graph[curr], word)
		}
	}
	return graph
}

type pair struct { word string; length int }

---

func ladderLength(beginWord, endWord string, wordList []string) int {
    q := list.New()
    q.PushBack(beginWord) 
    dist := map[string]int{beginWord: 1}
    for q.Len() != 0 {
        from := q.Front().Value.(string)
        q.Remove(q.Front())
        for _, to := range neighbors(from, wordList) {
            if _, ok := dist[to]; ok { continue }
            dist[to] = dist[from] + 1
            if to == endWord {
                return dist[to]
            }
            q.PushBack(to)
        }
    }
    return 0
}

func neighbors(from string, wordList []string) []string {
    var res []string
    for _, to := range wordList {
        if valid(from, to) { res = append(res, to) }
    }
    return res
}

func valid(from, to string) bool {
    if len(from) != len(to) { return false }
    var diff int
    for i := 0; i < len(from); i++ {
        if from[i] != to[i] {
            if diff == 1 { return false }
            diff++
        }
    }
    return diff == 1
}

---

// Sub-optimal solution. The graph construction takes O(n^2 * m)
// where n is the length of wordList and m is the length of each word
func ladderLength(begin string, end string, words []string) int {
    graph := newGraph(append(words, begin))
    q := list.New()
    q.PushBack(pair{word: begin, length: 1})
    visited := map[string]bool{begin: true}
    for q.Len() != 0 {
        curr := q.Front().Value.(pair)
        q.Remove(q.Front())
        for _, word := range graph[curr.word] {
            if word == end {
                return curr.length+1
            }
            if !visited[word] {
                visited[word] = true
                q.PushBack(pair{word: word, length: curr.length+1})
            }
        }
    }
    return 0
}

func newGraph(words []string) map[string][]string {
    graph := make(map[string][]string)
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            if i != j && validMove(words[i], words[j]) {
                graph[words[i]] = append(graph[words[i]], words[j])
            }
        }
    }
    return graph
}

func validMove(from, to string) bool {
    var diff int
    for i := 0; i < len(from); i++ {
        if from[i] != to[i] {
            if diff == 1 {
                return false
            }
            diff++
        }
    }
    return diff == 1
}

type pair struct {
    word string
    length int
}

---

func ladderLength(beginWord, endWord string, wordList []string) int {
    q := list.New()
    q.PushBack(beginWord)
    visited := map[string]bool{}
    dict := newDictionary(wordList)
    var level int = 1
    for q.Len() != 0 {
        size := q.Len()
        level++
        for i := 0; i < size; i++ {
            word := q.Front().Value.(string)
            q.Remove(q.Front())
            for _, neighbor := range getNeighbors(word) {
                if !dict[neighbor] || visited[neighbor] { continue }
                if neighbor == endWord { return level }
                visited[neighbor] = true
                q.PushBack(neighbor)
            }
        }
    }
    return 0
}

func newDictionary(words []string) map[string]bool {
    dict := make(map[string]bool)
    for _, word := range words {
        dict[word] = true
    }
    return dict
}

func getNeighbors(word string) []string {
    bslice := []byte(word)
    var neighbors []string
    for i := 0; i < len(word); i++ {
        orig := bslice[i]
        for c := byte('a'); c <= 'z'; c++ {
            if c == orig { continue }
            bslice[i] = c
            neighbors = append(neighbors, string(bslice))
        }
        bslice[i] = orig
    }
    return neighbors
}
