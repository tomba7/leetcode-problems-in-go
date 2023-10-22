type WordDictionary struct {
    root *TrieNode
}

type TrieNode struct {
    children map[byte]*TrieNode
    end bool
}

func Constructor() WordDictionary {
    return WordDictionary{
        root: &TrieNode{
            end: false, children: make(map[byte]*TrieNode),
        },
    }
}

func (d *WordDictionary) AddWord(word string)  {
    curr := d.root
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if _, exist := curr.children[ch]; !exist {
            curr.children[ch] = &TrieNode{
                children: make(map[byte]*TrieNode),
            }
        }
        curr = curr.children[ch]
    }
    curr.end = true
}

func (d *WordDictionary) Search(word string) bool {
    return search(word, d.root)
}

func search(word string, root *TrieNode) bool {
    curr := root
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if ch == '.' {
            for _, child := range curr.children {
                if search(word[i+1:], child) { return true }
            }
            return false
        } else {
            if _, exist := curr.children[ch]; !exist {
                return false
            }
            curr = curr.children[ch]
        }
    }
    return curr.end
}
