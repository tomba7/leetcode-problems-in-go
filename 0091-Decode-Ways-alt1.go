func numDecodings(s string) int {
    if len(s) == 0 {
        return 0
    } else if len(s) == 1 {
        if s[0] == '0' { return 0 }
        return 1
    }
    cache := make([]int, len(s) + 1)
    cache[0], cache[1] = 1, 1
    if s[0] == '0' { cache[1] = 0 }
    for i := 2; i < len(cache); i++ {
        cache[i] = cache[i - 1] * valid1(s[i - 1]) + cache[i - 2] * valid2(s[i - 2], s[i - 1])
    }
    return cache[len(s)]
}

func valid1(x byte) int {
    if x >= '1' && x <= '9' { return 1 }
    return 0
}

func valid2(x, y byte) int {
    if x == '1' || x == '2' && y <= '6' { return 1 }
    return 0
}
