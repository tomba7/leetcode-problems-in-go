func convert(s string, numRows int) string {
    if len(s) == 0 || numRows <= 1 { return s }

    buf := make([][]byte, numRows)
    result := []byte{}
    var down bool
    row := 0

    for i := 0; i < len(s); i++ {
        buf[row] = append(buf[row], s[i])
        if row == numRows - 1 { down = false }
        if row == 0 { down = true }
        if down { row++ } else { row-- }
    }

    for i := 0; i < len(buf); i++ {
        result = append(result, buf[i]...)
    }

    return string(result)
}
