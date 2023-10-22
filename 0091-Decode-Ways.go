func numDecodings(s string) int {
    if len(s) == 0 { return 0 }
    cache := make([]int, len(s))
    for i := range cache { cache[i] = -1 }
    return numDecodingsHelper(s, len(s) - 1, cache)
}

func numDecodingsHelper(s string, i int, cache []int) int {
    if i == -1 || i == 0 {
        if i == 0 && s[0] == '0' { return 0 };
        return 1
    }
    if cache[i] == -1 {
        cache [i] = numDecodingsHelper(s, i - 1, cache) * valid1(s[i]) +
                    numDecodingsHelper(s, i - 2, cache) * valid2(s[i - 1], s[i])
    }
    return cache[i]
}

func valid1(x byte) int {
    if x >= '1' && x <= '9' { return 1 }
    return 0
}

func valid2(x, y byte) int {
    if x == '1' || x == '2' && y <= '6' { return 1 }
    return 0
}

---

func numDecodings(s string) int {
    cache := make(map[string]int)
    return numDecodingsHelper(s, 0, cache)
}

func numDecodingsHelper(s string, i int, cache map[string]int) int {
    if val, exist := cache[s[i:]]; exist {
        return  val
    }
    n := len(s)
    if i == n {
        return 1
    }
    var res int
    if s[i] != '0' {
        res += numDecodingsHelper(s, i+1, cache)
        if i + 1 < n {
            if num, _ := strconv.Atoi(s[i:i+2]); num <= 26 {
                res += numDecodingsHelper(s, i+2, cache)
            }
        }
    }
    cache[s[i:]] = res
    return res
}
