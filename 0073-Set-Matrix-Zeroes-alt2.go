func setZeroes(matrix [][]int)  {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return
    }
    firstRowHasZero, firstColHasZero := false, false
    for i := 0; i < len(matrix); i++ {
        if matrix[i][0] == 0 {
            firstColHasZero = true
        }
    }
    for j := 0; j < len(matrix[0]); j++ {
        if matrix[0][j] == 0 {
            firstRowHasZero = true
        }
    }
    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i := len(matrix) - 1; i >= 1; i-- {
        for j := len(matrix[0]) - 1; j >= 1; j-- {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    if firstColHasZero {
        for i := 0; i < len(matrix); i++ {
            matrix[i][0] = 0
        }
    }
    if firstRowHasZero {
        for j := 0; j < len(matrix[0]); j++ {
            matrix[0][j] = 0
        }
    }
}
