var solution = func(read4 func([]byte) int) func([]byte, int) int {
    return func(buf []byte, n int) int {
        buf4 := make([]byte, 4)
        var i int
        for i < n {
            readCount := read4(buf4)
            for j := 0; j < readCount && i < n; i, j = i+1, j+1 {
                buf[i] = buf4[j]
            }
            if readCount != 4 { break }
        }
        return i
    }
}

---

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    return func(buf []byte, n int) int {
        buf4 := make([]byte, 4)
        var i int
        for i < n {
            m := read4(buf4)
            for j := 0; j < m && i < n; j, i = j+1, i+1 {
                buf[i] = buf4[j]
            }
            if m < 4 { break }
        }
        return i
    }
}

---

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    return func(buf []byte, n int) int {
        var i int
        buf4 := make([]byte, 4)
        for i < n {
            read := read4(buf4)
            for j := 0; j < read; j++ {
                buf[i] = buf4[j]
                i++
                if i == n { break }
            }
            if read != 4 { break }
        }
        return i
    }
}

---

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    return func(buf []byte, n int) int {
        i, numBytesRead := 0, 4
        buf4 := make([]byte, 4)
        for i < n && numBytesRead == 4 {
            numBytesRead = read4(buf4)
            for j := 0; j < numBytesRead; j++ {
                if i == n { return i }
                buf[i] = buf4[j]
                i++
            }
        }
        return i
    }
}
