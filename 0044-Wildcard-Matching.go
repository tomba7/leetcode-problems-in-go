/*
    s = "aa"
    p = "*****"
    Base Cases:
    i == m && j == n
        -> True
    i != m && j == n
        -> False
    i == m && j != n
        if any(p[j]) != '*'
            -> False
        else
            -> True
    General:
    s[i] == p[j]  or '?'
        -> increment both i and j
    '*'
        -> f(i, j+1) || f(i+1, j)
        
         *.  a.  *.  b.  -
         0.  1.  2.  3.  4   
       ---------------------
 a  0  |   |.  |.  |.  | F |
 d  1  |   |.  |.  |.  | F |
 c  2  |   |.  |.  |.  | F |
 e  3  |   |.  |.  |.  | F |
 b  4  |   |.  |.  |.  | F |
 -  5  | F | F | F | F | T |
       ---------------------
 */
func isMatch(s, p string) bool {
    m, n := len(s), len(p)
    dp := newDPTable(s, p)
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if s[i] == p[j] || p[j] == '?' {
                dp[i][j] = dp[i+1][j+1]
            } else if p[j] == '*' {
                dp[i][j] = dp[i+1][j] || dp[i][j+1]
            }
        }
    }
    return dp[0][0]
}

func newDPTable(s, p string) [][]bool {
    m, n := len(s), len(p)
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    dp[m][n] = true
    for j := n-1; j >= 0; j-- {
        if p[j] != '*' { break }
        dp[m][j] = true
    }
    return dp
}

---

func isMatch(s string, p string) bool {
    n, m := len(s), len(p)
    dp := newDPTable(s, p)
    for i := n-1; i >= 0; i-- {
        for j := m-1; j >= 0; j-- {
            if p[j] == '?' {
                dp[i][j] = dp[i+1][j+1]
            } else if p[j] == '*' {
                dp[i][j] = dp[i][j+1] || dp[i+1][j]
            } else {
                dp[i][j] = s[i] == p[j] && dp[i+1][j+1]
            }
        }
    }
    return dp[0][0]
}

func newDPTable(s, p string) [][]bool {
    n, m := len(s), len(p)
    dp := make([][]bool, n+1)
    for i := range dp {
        dp[i] = make([]bool, m+1)
    }
    dp[n][m] = true
    for i := 0; i < n; i++ {
        dp[i][m] = false
    }
    for j := m-1; j >= 0; j-- {
        if p[j] == '*' { dp[n][j] = dp[n][j+1] } else { break }
    }
    return dp
}

---

// TLE
func isMatch(s string, p string) bool {
    return helper(s, p, 0, 0)
}

func helper(s, p string, i, j int) bool {
    n, m := len(s), len(p)
    if i == n && j == m {
        return true
    } else if i != n && j == m {
        return false
    } else if i == n && j != m {
        for ; j < m; j++ {
            if p[j] != '*' { return false }
        }
        return true
    } else if p[j] == '?' {
        return helper(s, p, i+1, j+1)
    } else if p[j] == '*' {
        return helper(s, p, i, j+1) || helper(s, p, i+1, j)
    } else {
        return s[i] == p[j] && helper(s, p, i+1, j+1)
    }
}

---

// TLE
func isMatch(s string, p string) bool {
    return helper(s, p, 0, 0)
}

func helper(s, p string, i, j int) bool {
    m, n := len(s), len(p)
    if i == m && j == n {
        return true
    } else if i == m {
        for ; j < n; j++ {
            if p[j] != '*' { return false }
        }
        return true
    } else if j == n {
        return false
    }
    if s[i] == p[j] || p[j] == '?' {
        return helper(s, p, i+1, j+1)
    } else if p[j] == '*' {
        return helper(s, p, i, j+1) || helper(s, p, i+1, j)
    }
    return false
}

---

func isMatch(s, p string) bool {
    cache := map[string]bool{}
    return helper(s, p, 0, 0, cache)
}

func helper(s, p string, i, j int, cache map[string]bool) bool {
    if i == len(s) && j == len(p) {
        return true
    } else if j == len(p) {
        return false
    } else if i == len(s) {
        for ; j < len(p); j++ {
            if p[j] != '*' { return false }
        }
        return true
    }
    key := getKey(i, j)
    if match, found := cache[key]; found {
        return match
    }
    if p[j] == '?' || p[j] == s[i] {
        cache[key] = helper(s, p, i+1, j+1, cache)
    } else if p[j] == '*' {
        cache[key] = helper(s, p, i, j+1, cache) || helper(s, p, i+1, j, cache)
    }
    return cache[key]
}

func getKey(i, j int) string { return fmt.Sprintf("%d,%d", i, j) }

---

func isMatch(s, p string) bool {
    cache := map[string]bool{}
    return helper(s, p, 0, 0, cache)
}

func helper(s, p string, i, j int, cache map[string]bool) bool {
    if i == len(s) && j == len(p) {
        return true
    } else if j == len(p) {
        return false
    } else if i == len(s) {
        for ; j < len(p); j++ {
            if p[j] != '*' { return false }
        }
        return true
    }
    key := getKey(i, j)
    if match, found := cache[key]; found {
        return match
    }
    if p[j] == '?' || p[j] == s[i] {
            cache[key] = helper(s, p, i+1, j+1, cache)
    } else if p[j] == '*' {
        for ; i <= len(s); i++ {
            cache[key] = helper(s, p, i, j+1, cache)
            if cache[key] { break }
        }
    }
    return cache[key]
}

func getKey(i, j int) string { return fmt.Sprintf("%d,%d", i, j) }
