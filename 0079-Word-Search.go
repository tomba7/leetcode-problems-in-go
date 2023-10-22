func exist(board [][]byte, word string) bool {
    if len(word) == 0 || len(board) == 0 || len(board[0]) == 0 { return false }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if existHelper(board, word, i, j, 0) { return true }
        }
    }
    return false
}

func existHelper(board [][]byte, word string, row, col, indx int) bool {
    if indx == len(word) {
        return true
    } else if row < 0 || col < 0 || row == len(board) || col == len(board[0]) ||
        board[row][col] != word[indx] || board[row][col] == '*' {
        return false
    }
    indx++
    c := board[row][col]
    board[row][col] = '*'
    rc := existHelper(board, word, row + 1, col, indx) ||
          existHelper(board, word, row - 1, col, indx) ||
          existHelper(board, word, row, col + 1, indx) ||
          existHelper(board, word, row, col - 1, indx)
    board[row][col] = c
    return rc
}

---

/*
 - Start the DFS from every cell (if word[0] == board[i][j])
 - Before each DFS, create a seen matrix[m][n]
 - As we go deeper in the DFS increment i so we handle the next word[i+1]
 - If i == len(word) return true
 
 ["A","B","C","E"]
   
 ["S","F","C","S"]
               
 ["A","D","E","E"]
           ^
 */
func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    seen := newSeen(m, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] != word[0] { continue }
            if dfs(i, j, 0, board, word, seen) {
                return true
            }
        }
    }
    return false
}

func dfs(i, j, k int, board [][]byte, word string, seen [][]bool) bool {
    m, n := len(board), len(board[0])
    if k == len(word) { return true }
    if i < 0 || i == m || j < 0 || j == n || board[i][j] != word[k] || seen[i][j] {
        return false
    }
    seen[i][j] = true
    moves := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
    for _, move := range moves {
        row, col := i + move[0], j + move[1]
        if dfs(row, col, k+1, board, word, seen) {
            return true
        }
    }
    seen[i][j] = false
    return false
}

func newSeen(m, n int) [][]bool {
    seen := make([][]bool, m)
    for i := range seen {
        seen[i] = make([]bool, n)
    }
    return seen
}
