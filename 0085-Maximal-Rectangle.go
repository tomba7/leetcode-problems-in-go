// We form a histogram for every column
func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    if m == 0 {
        return 0
    }
    n := len(matrix[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    var res int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '0' {
                continue
            }
            if j == 0 {
                dp[i][j] = 1
            } else {
                dp[i][j] = 1 + dp[i][j-1]
            }
            width := dp[i][j]
            for k := i; k >= 0; k-- {
                width = min(width, dp[k][j])
                res = max(res, width * (i-k+1))
            }
        }
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }
func min(x, y int) int { if x < y { return x }; return y }

---

// Using the largest rectangular area in a histogram. We form
// a histogram for every row
func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    if m == 0 {
        return 0
    }
    n := len(matrix[0])
    dp := make([]int, n)
    var res int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                dp[j] += 1
            } else {
                dp[j] = 0
            }
        }
        res = max(res, largestRectangleArea(dp))
    }
    return res
}

func largestRectangleArea(heights []int) int {
    s := list.New()
    var res int
    n := len(heights)
    for i := 0; i <= n; i++ {
        for s.Len() != 0 && (i == n || heights[i] < heights[top(s)]) {
            j := s.Remove(s.Back()).(int)
            if s.Len() == 0 {
                res = max(res, heights[j] * i)
            } else {
                res = max(res, heights[j] * (i - top(s) - 1))
            }
        }
        if i < n {
            s.PushBack(i)
        }
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }
func top(s *list.List) int { return s.Back().Value.(int) }

---

// Brute force: O(m^3 * n^3)
func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    if m == 0 { return 0 }
    n := len(matrix[0])
    var res int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '0' { continue }
            res = max(res, maxArea(matrix, i, j))
        }
    }
    return res
}

func maxArea(matrix [][]byte, r, c int) int {
    m, n := len(matrix), len(matrix[0])
    var res int
    for i := r; i < m; i++ {
        for j := c; j < n; j++ {
            res = max(res, area(matrix, r, c, i, j))
        }
    }
    return res
}

func area(matrix [][]byte, startRow, startCol, endRow, endCol int) int {
    allOnes := true
    for i := startRow; i <= endRow; i++ {
        for j := startCol; j <= endCol; j++ {
            if matrix[i][j] == '0' {
                allOnes = false
            }
        }
    }
    if allOnes {
        return (endRow - startRow + 1) * (endCol - startCol + 1)
    }
    return 0
}

func max(x, y int) int { if x > y { return x }; return y }
