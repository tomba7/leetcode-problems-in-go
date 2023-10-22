func leftMostColumnWithOne(matrix BinaryMatrix) int {
    dims := matrix.Dimensions()
    m, n := dims[0], dims[1]
    col := -1
    for r, c := 0, n-1; r < m && c >= 0; {
        if matrix.Get(r, c) == 1 {
            col = c
            c--
        } else {
            r++
        }
    }
    return col
}

---

func leftMostColumnWithOne(matrix BinaryMatrix) int {
    dimensions := matrix.Dimensions() 
    m, n := dimensions[0], dimensions[1]
    i, j := 0, n-1
    for i < m && j >= 0 {
        if matrix.Get(i, j) == 1 {
            j--
        } else {
            i++
        }
    }
    if j == n-1 { return -1 }
    return j+1
}
