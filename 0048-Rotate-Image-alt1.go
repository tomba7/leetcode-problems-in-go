func rotate(matrix [][]int)  {
    for len(matrix) == 0 || len(matrix) != len(matrix[0]) { return }
    n := len(matrix)
    for offset := 0; offset < n / 2; offset++ {
        for i := offset; i < n - offset - 1; i++ {
            j := n - i - 1
            x := matrix[offset][i]
            matrix[offset][i] = matrix[j][offset]
            matrix[j][offset] = matrix[n - offset - 1][j]
            matrix[n - offset - 1][j] = matrix[i][n - offset - 1]
            matrix[i][n - offset - 1] = x
        }
    }
}
