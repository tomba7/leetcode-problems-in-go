func numDecodings(s string) int {
    if len(s) == 0 { return 0 }
    cache := make([]int, len(s) + 1) 
    cache[0], cache[1] = 1, 1
    if s[0] == '0' {
        cache[1] = 0
    } else if s[0] == '*' {
        cache[1] = 9
    }
    for i := 2; i < len(cache); i++ {
        cache[i] = cache[i - 1] * valid1(s[i - 1]) + cache[i - 2] * valid2(s[i - 2], s[i - 1])
        cache[i] %= 1000000007
    }
    return cache[len(s)]
}

func valid1(x byte) int {
    if x == '*' { return 9 } else if x >= '1' && x <= '9' { return 1 }
    return 0
}

func valid2(x, y byte) int {
    if x == '*' && y == '*' {
        return 15  // (11, 12, 13 ... 19, 21, 22 ... 26)
    } else if x == '*' {
        if y >= '0' && y <= '6' { return 2 }
        return 1
    } else if y == '*' {
        if x == '1' { return 9 } else if x == '2' { return 6 }
        return 0
    }
    if x == '1' || x == '2' && y <= '6' { return 1 }
    return 0
}
