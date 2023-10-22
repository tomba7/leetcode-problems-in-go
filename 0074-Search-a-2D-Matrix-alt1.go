func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return false }
    lo, hi := 0, len(matrix) * len(matrix[0]) - 1
    for lo < hi {
        mid := lo + (hi - lo)/2
        if target > getValue(matrix, mid) {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    return getValue(matrix, lo) == target
}

func getValue(matrix [][]int, index int) int {
    i, j := index / len(matrix[0]), index % len(matrix[0])
    return matrix[i][j]
}
