/*
    s = "daab"
    p = "ec*a*b"

  Base Cases:
    i == m && j == n:
        True
    j == n:
        False
    i == m:
        if j+1 < n && p[j+1] == '*':
            return f(i, j+2)
        else
            False
        
  General Cases:
    j+1 < n && p[j+1] == '*':
        f(i, j+2) || (s[i] == p[j] || p[j] == '.') && f(i+1, j)
    s[i] == p[j] || p[j] == '.':
        f(i+1, j+1)
    any:
        False
 */

func isMatch(s string, p string) bool {
    return helper(s, p, 0, 0)
}

func helper(s, p string, i, j int) bool {
    m, n := len(s), len(p)
    if i == m && j == n {
        return true
    } else if j == n {
        return false
    } else if i == m {
        if j+1 < n && p[j+1] == '*' { return helper(s, p, i, j+2) }
        return false
    }
    if j+1 < n && p[j+1] == '*' {
        return helper(s, p, i, j+2) || (s[i] == p[j] || p[j] == '.') && helper(s, p, i+1, j)
    } else if s[i] == p[j] || p[j] == '.' {
        return helper(s, p, i+1, j+1)
    }
    return false
}

---

func isMatch(s string, p string) bool {
    return helper(s, p, 0, 0)
}

func helper(s, p string, i, j int) bool {
    m, n := len(s), len(p)
    if j == n { return i == m }
    if j+1 < n && p[j+1] == '*' {
        return helper(s, p, i, j+2) ||
               i < m && (s[i] == p[j] || p[j] == '.') && helper(s, p, i+1, j)
    } else if i < m && s[i] == p[j] || p[j] == '.' {
        return helper(s, p, i+1, j+1)
    }
    return false
}

---

/*
 s = "aab", p = "c*a*b"
 m = 3, n = 5
 
           c.  *.  a.  *.  b.  -
           0.  1.  2.  3.  4.  5
         +-----------------------+
   a  0  | . | . | . | . | . | F |
   a  1  | . | . | . | . | . | F |
   b  2  | . | . | . | . | . | F |
   -  3  | . | . | . | . | . | T |
         +-----------------------+
 */
func isMatch(s string, p string) bool {
    dp := newDPTable(s, p)
    m, n := len(s), len(p)
    for i := m; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            match := i < m && (s[i] == p[j] || p[j] == '.') 
            if j+1 < n && p[j+1] == '*' {
                dp[i][j] =  dp[i][j+2] || match && dp[i+1][j]
            } else if match {
                dp[i][j] =  dp[i+1][j+1]
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
    return dp
}

---

func isMatch(s, p string) bool {
    n, m := len(s), len(p)
    dp := newDPTable(s, p)
    for i := n; i >= 0; i-- {
        for j := m-1; j >= 0; j-- {
            match := i != n && (s[i] == p[j] || p[j] == '.')
            if j+1 < m && p[j+1] == '*' {
                dp[i][j] = dp[i][j+2] || match && dp[i+1][j]
            } else {
                dp[i][j] = match && dp[i+1][j+1]
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
    return dp
}

---

func isMatch(s string, p string) bool {
    cache := map[string]bool{}
    return helper(s, p, 0, 0, cache)
}

func helper(s, p string, i, j int, cache map[string]bool) bool {
    n, m := len(s), len(p)
    if j == m { return i == n }
    key := fmt.Sprintf("%d,%d", i, j)
    if res, found := cache[key]; found {
        return res
    }
    match := i != n && (s[i] == p[j] || p[j] == '.')
    if j+1 < m && p[j+1] == '*' {
        cache[key] = helper(s, p, i, j+2, cache) || match && helper(s, p, i+1, j, cache)
    } else {
        cache[key] = match && helper(s, p, i+1, j+1, cache)
    }
    return cache[key]
}

---

func isMatch(s string, p string) bool {
    cache := map[string]bool{}
    return helper(s, p, 0, 0, cache)
}

func helper(s, p string, i, j int, cache map[string]bool) bool {
    n, m := len(s), len(p)
    if j == m {
        return i == n
    } else if i == n && j != m {
        // this clause can be condensed/combined in the main section below
        if j+1 != m && p[j+1] == '*' {
            return helper(s, p, i, j+2, cache)
        } else {
            return false
        }
    }
    key := fmt.Sprintf("%d,%d", i, j)
    if match, found := cache[key]; found {
        return match
    }
    match := s[i] == p[j] || p[j] == '.'
    if j+1 != m && p[j+1] == '*' {
        cache[key] = helper(s, p, i, j+2, cache) || match && helper(s, p, i+1, j, cache)
    } else {
        cache[key] = match && helper(s, p, i+1, j+1, cache)
    }
    return cache[key]
}

---

// Convoluted first pass solution (but works)
func isMatch(s string, p string) bool {
    cache := map[string]bool{}
    return helper(s, p, 0, 0, cache)
}

func helper(s, p string, i, j int, cache map[string]bool) bool {
    if j == len(p) { return i == len(s) }
    key := fmt.Sprintf("%d,%d", i, j)
    if match, found := cache[key]; found { return match }
    if j+1 < len(p) && p[j+1] == '*' {
        if i != len(s) && (s[i] == p[j] || p[j] == '.') {
            cache[key] = helper(s, p, i+1, j, cache) || helper(s, p, i, j+2, cache)
        } else {
            cache[key] = helper(s, p, i, j+2, cache)
        }
    } else {
        if i != len(s) && (s[i] == p[j] || p[j] == '.') {
            cache[key] = helper(s, p, i+1, j+1, cache)
        }
    }
    return cache[key]
}
