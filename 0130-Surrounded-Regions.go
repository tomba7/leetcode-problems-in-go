func solve(board [][]byte)  {
    // Ignore any board sizes that are less than 2 x 2
    if len(board) == 0 || len(board[0]) == 0 ||  len(board) <= 2 || len(board[0]) <= 2 {return}
    // Walk the leftmost and rightmost columns. Look for any 'O's
    // and perform DFS to replace neighboring 'O's with '-'s
    for i := 0; i < len(board); i++ {
        visitNeighbors(board, i, 0)
        visitNeighbors(board, i, len(board[0]) - 1)
    }
    // Walk the top and bottom rows. Perform the same DFS
    // here to replace neighboring 'O's with '-'s
    for j := 0; j < len(board[0]); j++ {
        visitNeighbors(board, 0, j)
        visitNeighbors(board, len(board) - 1, j)
    }
    // At this point, any unbounded 'O's are marked as '-'. Mark them back as 'O's
    // On the other hand, all bounded 'O's remains untouched. Mark them now as 'X's
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            } else if board[i][j] == '-' {
                board[i][j] = 'O'
            }
        }
    }
}

func visitNeighbors(board [][]byte, i, j int) {
    // Ignore DFS, if any of the co-ordinates are out of bounds.
    // Or if the node/cell is not 'O'.
    for i < 0 || j < 0 || i > len(board) - 1 || j > len(board[0]) - 1 || board[i][j] != 'O' {
        return
    }
    // Replace any unbounded 'O' with '-'
    board[i][j] = '-'
    visitNeighbors(board, i - 1, j)
    visitNeighbors(board, i, j + 1)
    visitNeighbors(board, i + 1, j)
    visitNeighbors(board, i, j - 1)
}

---

func solve(board [][]byte)  {
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        flagBorderCells(board, i, 0)
        flagBorderCells(board, i, n-1)
    }
    for j := 1; j < n-1; j++ {
        flagBorderCells(board, 0, j)
        flagBorderCells(board, m-1, j)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'B' {
                board[i][j] = 'O'
            } else if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
        }
    }
}

func flagBorderCells(board [][]byte, i, j int) {
    m, n := len(board), len(board[0])
    if i < 0 || i == m || j < 0 || j == n || board[i][j] != 'O' {
        return
    }
    board[i][j] = 'B'
    offsets := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    for _, offset := range offsets {
        r, c := i + offset[0], j + offset[1]
        flagBorderCells(board, r, c)
    }
}
