func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return false }
    for i, j := 0, len(matrix[0]) - 1; i < len(matrix) && j >= 0; {
        if matrix[i][j] < target {
            i++
        } else if matrix[i][j] > target {
            j--
        } else {
            return true
        }
    }
    return false
}
