func rotate(matrix [][]int)  {
    if len(matrix) == 0 || len(matrix) != len(matrix[0]) { return }
    for layer := 0; layer < len(matrix)/2; layer++ {
        first, last := layer, len(matrix) - 1 - layer
        for i, j := first, last; i < last; i, j = i+1, j-1 {
            topLeft := matrix[first][i]
            matrix[first][i] = matrix[j][first]
            matrix[j][first] = matrix[last][j]
            matrix[last][j] = matrix[i][last]
            matrix[i][last] = topLeft
        }
    }
}
