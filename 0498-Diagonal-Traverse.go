func findDiagonalOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return nil }
    dirMap := [][]int{{-1, 1}, {1, -1}}
    M, N := len(matrix), len(matrix[0])
    result := make([]int, M*N)
    var m, n, d int
    for i := 0; i < M*N; i++ {
        result[i] = matrix[m][n]
        m, n = m + dirMap[d][0], n + dirMap[d][1]
        if m == M {
            m, n, d = m-1, n+2, 1-d
        } else if n == N {
            m, n, d = m+2, n-1, 1-d
        } else if m < 0 {
            m, d = 0, 1-d
        } else if n < 0 {
            n, d = 0, 1-d
        }
    }
    return result
}
