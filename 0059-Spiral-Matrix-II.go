func generateMatrix(n int) [][]int {
    var numIndex int
    matrix := make([][]int, n)
    for i := 0; i < n; i++ { matrix[i] = make([]int, n) }
    for i := 0; i < (n + 1)/2; i++ {
        if i == n-i-1 {
            matrix[i][i] = numIndex + 1
            break
        }
        for j := i; j < n-i-1; j++ {
            matrix[i][j] = numIndex + 1
            numIndex++
        }
        for j := i; j < n-i-1; j++ {
            matrix[j][n-i-1] = numIndex + 1
            numIndex++
        }
        for j := n-i-1; j > i; j-- {
            matrix[n-i-1][j] = numIndex + 1
            numIndex++
        }
        for j := n-i-1; j > i; j-- {
            matrix[j][i] = numIndex + 1
            numIndex++
        }
    }
    return matrix
}
