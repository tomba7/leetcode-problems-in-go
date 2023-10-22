var solution = func(read4 func([]byte) int) func([]byte, int) int {
    buf4 := make([]byte, 4)
    j, numRead := 4, 0
    return func(buf []byte, n int) int {
        i := 0
        for i < n {
            if j == 4 {
                numRead = read4(buf4)
                j = 0
            }
            for ; j < numRead && i < n; i, j = i+1, j+1 {
                buf[i] = buf4[j]
            }
            if j == numRead && numRead < 4 { break }
        }
        return i
    }
}

---

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    var j, read int
    buf4 := make([]byte, 4)
    return func(buf []byte, n int) int {
        var i int
        for ; i < n; i++ {
            if j == read {
                j = 0
                read = read4(buf4)
                if read == 0 {
                    return i
                }
            }
            buf[i] = buf4[j]
            j++
        }
        return i
    }
}

---

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    var j int
    read := 4
    buf4 := make([]byte, 4)
    return func(buf []byte, n int) int {
        var i int
        for i < n {
            for j < read {
                if j == 0 {
                    read = read4(buf4)
                    if read == 0 { return i }
                }
                buf[i] = buf4[j]
                i++
                j = (j+1)%4
                if i == n { return i }
            }
            if read != 4 { return i }
        }
        return i
    }
}
