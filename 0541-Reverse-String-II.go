func reverseStr(s string, k int) string {
    bs := []byte(s)
    for i := 0; i < len(s); i = i + 2 * k {
        reverse(bs, i, min(i + k - 1, len(s) - 1))
    }
    return string(bs)
}

func reverse(bs []byte, i, j int) {
    for ; i < j; i, j = i + 1, j - 1 {
        bs[i], bs[j] = bs[j], bs[i]
    }
}

func min(x, y int) int { if x < y { return x }; return y }
