func searchMatrix(arr [][]int, x int) bool {
    r, c := len(arr), len(arr[0])
    i, j := 0, c-1
    for i < r && j >= 0 {
        if x < arr[i][j] {
            j--
        } else if x > arr[i][j] {
            i++
        } else {
            return true
        }
    }
    return false
}

---

// Bad solution
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return false }
    lo, hi := 0, len(matrix) * len(matrix[0]) - 1
    for lo <= hi {
        mid := lo + (hi - lo)/2
        i, j := mid / len(matrix[0]), mid % len(matrix[0])
        if target < matrix[i][j] {
            hi = mid - 1
        } else if target > matrix[i][j] {
            lo = mid + 1
        } else {
            return true
        }
    }
    return false
}
