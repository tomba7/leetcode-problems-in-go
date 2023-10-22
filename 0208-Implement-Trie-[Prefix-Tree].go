type Trie struct {
    root *TrieNode
}

type TrieNode struct {
    endOfWord bool
    children map[byte]*TrieNode
}

func NewTrieNode() *TrieNode {
    return &TrieNode{children: make(map[byte]*TrieNode)}
}

func Constructor() Trie {
    return Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string)  {
    curr := t.root
    for i := 0; i < len(word); i++ {
        child, contains := curr.children[word[i]]
        if !contains {
            child = NewTrieNode()
            curr.children[word[i]] = child
        }
        curr = child
    }
    curr.endOfWord = true
}

func (t *Trie) Search(word string) bool {
    curr := t.root
    for i := 0; i < len(word); i++ {
        child, contains := curr.children[word[i]]
        if !contains { return false }
        curr = child
    }
    return curr.endOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
    curr := t.root
    for i := 0; i < len(prefix); i++ {
        child, contains := curr.children[prefix[i]]
        if !contains { return false }
        curr = child
    }
    return true
}

---

type Trie struct {
    end bool
    children map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        end: false, children: make(map[byte]*Trie),
    }
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
    curr := t
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if _, exist := curr.children[ch]; !exist {
            newNode := Constructor()
            curr.children[ch] = &newNode
        }
        curr = curr.children[ch]
    }
    curr.end = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
    curr := t
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if _, exist := curr.children[ch]; !exist {
            return false
        }
        curr = curr.children[ch]
    }
    return curr.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
    curr := t
    for i := 0; i < len(prefix); i++ {
        ch := prefix[i]
        if _, exist := curr.children[ch]; !exist {
            return false
        }
        curr = curr.children[ch]
    }
    return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
