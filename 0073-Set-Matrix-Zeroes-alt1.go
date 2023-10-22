func setZeroes(matrix [][]int)  {
    if len(matrix) == 0 || len(matrix[0]) == 0 {return}
    rows := make([]bool, len(matrix))
    cols := make([]bool, len(matrix[0]))
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                rows[i], cols[j] = true, true
            }
        }
    }
    for row, zero := range rows {
        if zero == false {continue}
        for j := 0; j < len(matrix[0]); j++ {
            matrix[row][j] = 0
        }
    }
    for col, zero := range cols {
        if zero == false {continue}
        for i := 0; i < len(matrix); i++ {
            matrix[i][col] = 0
        }
    }
}
