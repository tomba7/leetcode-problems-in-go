func knightProbability(n, k, row, col int) float64 {
    dp := newDP(n, k)
    dp[0][row][col] = 1
    for i := 1; i <= k; i++ {
        for r := 0; r < n; r++ {
            for c := 0; c < n; c++ {
                for _, offset := range offsets {
                    nr, nc := r + offset[0], c + offset[1]
                    if nr < 0 || nr >= n || nc < 0 || nc >= n { continue }
                    dp[i][r][c] += dp[i-1][nr][nc]/8.0
                }
            }
        }
    }
    var res float64
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            res += dp[k][i][j]
        }
    }
    return res
}

func newDP(n, k int) [][][]float64 {
    dp := make([][][]float64, k+1)
    for i := range dp {
        dp[i] = make([][]float64, n)
        for j := range dp[i] {
            dp[i][j] = make([]float64, n)
        }
    }
    return dp
}

var offsets = [][]int{
    {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1},
}

---

func knightProbability(n, k, row, col int) float64 {
    dp := newDP(n, k)

    dp[0][row][col] = 1
    for i := 1; i <= k; i++ {
        for r := 0; r < n; r++ {
            for c := 0; c < n; c++ {
                for _, offset := range offsets {
                    nr, nc := r + offset[0], c + offset[1]
                    if nr < 0 || nr >= n || nc < 0 || nc >= n { continue }
                    dp[k][r][c] += dp[k-1][nr][nc]*1.0/8.0
                }
            }
        }
    }
    var res float64
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            res += dp[k][i][j]
        }
    }
    return res
}

func newDP(n, k int) [][]float64 {
    dp := make([][][]float64, k+1)
    for i := range dp {
        dp[i] = make([][]float64, n)
        for j := range dp[i] {
            dp[i][j] = make([]float64, n)
        }
    }
    return dp
}

var offsets = [][]int{
    {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1},
}

---

func knightProbability(n, k, row, col int) float64 {
    dp := newDP(n)
    dp[row][col] = 1
    for ; k > 0; k-- {
        dp2 := newDP(n)
        for r := 0; r < n; r++ {
            for c := 0; c < n; c++ {
                for k := 0; k < 8; k++ {
                    nr, nc := r + offsets[k][0], c + offsets[k][1]
                    if nr < 0 || nr >= n || nc < 0 || nc >= n { continue }
                    dp2[nr][nc] += dp[r][c]/8.0
                }
            }
        }
        dp = dp2
    }
    res := 0.0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            res += dp[i][j]
        }
    }
    return res
}

func newDP(n int) [][]float64 {
    dp := make([][]float64, n)
    for i := range dp { dp[i] = make([]float64, n) }
    return dp
}

var offsets = [][]int{
    {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1},
}

---

/*  Recursion (Time Limit Exceeded)

    0. 1. 2
    -- -- --
 0 |* |. |. |
 1 |. |. |a |
 2 |. |b |. |
    -- -- --
                       0,0                     k = 2    prob = 1/8
                       
                    /       \
                  
                1,2            2,1             k = 1    prob = 1/8
              
             /       \       /      \
             
          0,0         2,0  0,0       0,2       k = 0    prob = 1.0
 */

func knightProbability(n, k, row, col int) float64 {
    if row < 0 || row >= n || col < 0 || col >= n {
        return 0.0
    } else if k == 0 {
        return 1.0
    }
    var sum float64
    for _, offset := range offsets {
        r, c := row + offset[0], col + offset[1] 
        sum += 1.0/8.0 * knightProbability(n, k-1, r, c)
    }
    return sum
}

var offsets = [][]int{
    {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1},
}
