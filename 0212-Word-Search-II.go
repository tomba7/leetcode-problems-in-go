func findWords(board [][]byte, words []string) []string {
    result, root := []string{}, buildTrie(words)
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            findWordsHelper(board, i, j, root, &result)
        }
    }
    return result
}

func findWordsHelper(board [][]byte, i, j int, node *TrieNode, result *[]string) {
    if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || board[i][j] == '*' { return }
    c := board[i][j]
    child, ok := node.children[c]
    if !ok {
        return
    } else if child.word != "" {
        *result = append(*result, child.word)
        child.word = ""
    }
    board[i][j] = '*'
    findWordsHelper(board, i - 1, j, child, result)
    findWordsHelper(board, i, j - 1, child, result)
    findWordsHelper(board, i + 1, j, child, result)
    findWordsHelper(board, i, j + 1, child, result)
    board[i][j] = c
}

func buildTrie(words []string) *TrieNode {
    root := &TrieNode{children: map[byte]*TrieNode{}}
    for _, word := range words {
        node := root
        for i := 0; i < len(word); i++ {
            child, ok := node.children[word[i]]
            if !ok {
                child = &TrieNode{children: map[byte]*TrieNode{}}
                node.children[word[i]] = child
            }
            node = child
        }
        node.word = word
    }
    return root
}

type TrieNode struct {
    children map[byte]*TrieNode
    word string
}
