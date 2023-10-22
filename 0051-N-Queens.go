func solveNQueens(n int) [][]string {
    var res [][]string
    helper(make([]int, n), 0, &res)
    return res
}

func helper(queens []int, row int, res *[][]string) {
    n := len(queens)
    if row == n {
        *res = append(*res, encode(queens))
        return
    }
    for col := 0; col < n; col++ {
        queens[row] = col
        if !valid(queens, row) { continue  }
        helper(queens, row+1, res)
    }
}

func valid(queens []int, row int) bool {
    for r := 0; r < row; r++ {
        if queens[r] == queens[row] || row - r == abs(queens[row] - queens[r]) {
            return false
        }
    }
    return true
}

func encode(queens []int) []string {
    var res []string
    n := len(queens)
    for i := 0; i < n; i++ {
        bs := bytes.Repeat([]byte{'.'}, n)
        bs[queens[i]] = 'Q'
        res = append(res, string(bs))
    }
    return res
}

func abs(x int) int {
    if x < 0 { return -x }
    return x
}
