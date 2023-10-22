func suggestedProducts(products []string, searchWord string) [][]string {
    n := len(products)
    sort.Strings(products)
    var prefix string
    var res [][]string
    for _, ch := range searchWord {
        prefix += string(ch)
        j := sort.SearchStrings(products, prefix)
        var curr []string
        for i := 0; i < 3 && i + j < n; i++ {
            if strings.HasPrefix(products[i+j], prefix) {
                curr = append(curr, products[i+j])
            }
        }
        res = append(res, curr)
    }
    return res
}

---

func suggestedProducts(products []string, searchWord string) [][]string {
    root := &TrieNode{}
    sort.Strings(products)
    for _, product := range products {
        curr := root
        for _, ch := range product {
            if curr.children[ch-'a'] == nil {
                curr.children[ch-'a'] = &TrieNode{}
            }
            curr = curr.children[ch-'a']
            if len(curr.suggested) < 3 {
                curr.suggested = append(curr.suggested, product)
            }
        }
    }
    var res [][]string
    for _, ch := range searchWord {
        if root != nil {
            root = root.children[ch-'a']
        }
        if root == nil {
            res = append(res, []string{})
        } else {
            res = append(res, root.suggested)
        }
    }
    return res
}

type TrieNode struct {
    children [26]*TrieNode
    suggested []string
}

---

/*
 ["mobile", "moneypot", "monitor", "mouse", "mousepad"]

 m = ["mobile", "moneypot", "monitor"]
 o = ["mobile", "moneypot", "monitor"]
 u = ["mouse", "mousepad"]
 s = ["mouse", "mousepad"]
 e = ["mouse", "mousepad"]
 
 - Sort the products array
 - Insert each product into a trie
 - As we insert (in lexicographically sorted order), we store the
   first 3 strings encountered so far.
 - Finally, we take iterate the searchWord, at pick the strings
   stored at each index in the trie corresponding to the searchWord
 */
func suggestedProducts(products []string, searchWord string) [][]string {
    root := &TrieNode{children: make(map[byte]*TrieNode)}
    sort.Strings(products)
    for _, product := range products {
        curr := root
        for i := 0; i < len(product); i++ {
            ch := product[i]
            if _, contains := curr.children[ch]; !contains {
                curr.children[ch] = &TrieNode{children: make(map[byte]*TrieNode)}
            }
            curr = curr.children[ch]
            if len(curr.suggestions) < 3 {
                curr.suggestions = append(curr.suggestions, product)
            }
        }
    }
    var res [][]string
    curr := root
    for i := 0; i < len(searchWord); i++ {
        ch := searchWord[i]
        if curr != nil && curr.children[ch] != nil {
            curr = curr.children[ch]
            res = append(res, curr.suggestions)
        } else {
            curr = nil
            res = append(res, nil)
        }
    }
    return res
}

type TrieNode struct {
    children map[byte]*TrieNode
    suggestions []string
}
