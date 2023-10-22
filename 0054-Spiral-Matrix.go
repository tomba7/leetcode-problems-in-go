func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return nil }
    var result []int
    M, N := len(matrix), len(matrix[0])
    numLayers := (min(M, N)+1)/2
    for i := 0; i < numLayers; i++ {
        if i == M-i-1 {
            for n := i; n < N-i; n++ {
                result = append(result, matrix[i][n])
            }
            break
        } else if i == N-i-1 {
            for m := i; m < M-i; m++ {
                result = append(result, matrix[m][i])
            }
            break
        }
        for n := i; n < N-i-1; n++ { result = append(result, matrix[i][n]) }
        for m := i; m < M-i-1; m++ { result = append(result, matrix[m][N-i-1]) }
        for n := N-i-1; n > i; n-- { result = append(result, matrix[M-i-1][n]) }
        for m := M-i-1; m > i; m-- { result = append(result, matrix[m][i]) }
    }
    return result
}

func min(x, y int) int { if x < y { return x }; return y }

/*
 * Slightly different
 */
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return nil }
    M, N := len(matrix), len(matrix[0])
    numLayers := (min(M, N)+1)/2
    var result []int
    for i := 0; i < numLayers; i++ {
        if i == M-1-i {
            for n := i; n < N-i; n++ { result = append(result, matrix[i][n]) }
        } else if i == N-1-i {
            for m := i; m < M-i; m++ { result = append(result, matrix[m][i]) }
        } else {
            for n := i; n < N-1-i; n++ { result = append(result, matrix[i][n]) }
            for m := i; m < M-1-i; m++ { result = append(result, matrix[m][N-1-i]) }
            for n := N-1-i; n > i; n-- { result = append(result, matrix[M-1-i][n]) }
            for m := M-1-i; m > i; m-- { result = append(result, matrix[m][i]) }
        }
    }
    return result
}

func min(x, y int) int { if x < y { return x }; return y }
