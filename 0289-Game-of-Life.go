/*
    0 1 0   0 0 0
    0 0 1   1 0 1
    1 1 1   0 1 1
    0 0 0   0 1 0
    
    lc < 2 live      => dies
    lc == 2 | 3 live => lives
    lc > 3 live      => dies
    dc == 3 live     => lives

    Time: O(MN), Space(MN)
 */
func gameOfLife(board [][]int)  {
    m, n := len(board), len(board[0])
    duplicate := duplicateBoard(board)
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            liveNeighbors := getLiveNeighbors(duplicate, r, c)
            if duplicate[r][c] == 1 {
                if liveNeighbors < 2 || liveNeighbors > 3 {
                    board[r][c] = 0
                }
            } else {
                if liveNeighbors == 3 {
                    board[r][c] = 1
                }    
            }
        }
    }
}

func getLiveNeighbors(duplicate [][]int, row, col int) int {
    m, n := len(duplicate), len(duplicate[0])
    offset := []int{-1, 0, 1}
    var liveNeighbors int
    for i := 0; i < len(offset); i++ {
        for j := 0; j < len(offset); j++ {
            if i == 1 && j == 1 { continue }
            nextRow, nextCol := row + offset[i], col + offset[j]
            if nextRow < 0 || nextRow == m || nextCol < 0 || nextCol == n || duplicate[nextRow][nextCol] != 1 {
                continue
            }
            liveNeighbors++ 
        }
    }
    return liveNeighbors
}

func duplicateBoard(board [][]int) [][]int {
    m, n := len(board), len(board[0])
    duplicate := make([][]int, m)
    for i := range duplicate {
        duplicate[i] = make([]int, n)
        copy(duplicate[i], board[i])
    }
    return duplicate 
}

---

// Time: O(MN), Space(1)
func gameOfLife(board [][]int)  {
    m, n := len(board), len(board[0])
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            liveNeighbors := getLiveNeighbors(board, r, c)
            if board[r][c] == 1 {
                if liveNeighbors < 2 || liveNeighbors > 3 {
                    board[r][c] = -1
                }
            } else {
                if liveNeighbors == 3 {
                    board[r][c] = 2
                }
            }
        }
    }
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if board[r][c] > 0 {
                board[r][c] = 1
            } else {
                board[r][c] = 0
            }
        }
    }
}

func getLiveNeighbors(board [][]int, row, col int) int {
    m, n := len(board), len(board[0])
    offset := []int{-1, 0, 1}
    var liveNeighbors int
    for i := 0; i < len(offset); i++ {
        for j := 0; j < len(offset); j++ {
            if i == 1 && j == 1 { continue }
            nextRow, nextCol := row + offset[i], col + offset[j]
            if nextRow < 0 || nextRow == m || nextCol < 0 || nextCol == n || abs(board[nextRow][nextCol]) != 1 {
                continue
            }
            liveNeighbors++ 
        }
    }
    return liveNeighbors
}

func duplicateBoard(board [][]int) [][]int {
    m, n := len(board), len(board[0])
    duplicate := make([][]int, m)
    for i := range duplicate {
        duplicate[i] = make([]int, n)
        copy(duplicate[i], board[i])
    }
    return duplicate 
}

func abs(x int) int { if x < 0 { return -x }; return x }
