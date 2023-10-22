type StreamChecker struct {
    root *TrieNode
    stream *list.List
}

type TrieNode struct {
    children map[byte]*TrieNode
    word bool
}

func NewTrieNode() *TrieNode {
    return &TrieNode{children: make(map[byte]*TrieNode)}
}

func Constructor(words []string) StreamChecker {
    root := NewTrieNode()
    for i := 0; i < len(words); i++ {
        word := reverse(words[i])
        curr := root
        for j := 0; j < len(word); j++ {
            ch := word[j]
            child, contains := curr.children[ch]
            if !contains {
                child = NewTrieNode()
                curr.children[ch] = child
            }
            curr = child
        }
        curr.word = true
    }
    return StreamChecker{
        root: root, stream: list.New(),
    }
}

func (s *StreamChecker) Query(letter byte) bool {
    s.stream.PushBack(letter)
    curr := s.root
    for e := s.stream.Back(); e != nil; e = e.Prev() {
        ch := e.Value.(byte)
        if curr.word {
            return true
        }
        if _, contains := curr.children[ch]; !contains {
            return false
        }
        curr = curr.children[ch]
    }
    return curr.word
}

func reverse(s string) string {
    buf := []byte(s)
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        buf[i], buf[j] = buf[j], buf[i]
    }
    return string(buf)
}
