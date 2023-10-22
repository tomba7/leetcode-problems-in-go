func removeInvalidParentheses(s string) []string {
    extraLeft, extraRight := getExtraParentheses(s)
    var res []string
    removeParentheses(s, 0, 0, extraLeft, extraRight, '#', nil, &res)
    return res
}

func removeParentheses(s string, i, openCount, extraLeft, extraRight int,
                       skipped byte, path []byte, res *[]string) {
    if openCount < 0 || extraLeft < 0 || extraRight < 0 {
        return
    } else if i == len(s) {
        if openCount == 0 {
            *res = append(*res, string(path))
        }
        return
    }
    if s[i] == '(' {
        removeParentheses(s, i+1, openCount, extraLeft-1, extraRight, '(', path, res)
        path = append(path, s[i])
        if skipped != '(' {
            removeParentheses(s, i+1, openCount+1, extraLeft, extraRight, '#', path, res)
        }
    } else if s[i] == ')' {
        removeParentheses(s, i+1, openCount, extraLeft, extraRight-1, '#', path, res)
        path = append(path, s[i])
        if skipped != ')' {
            removeParentheses(s, i+1, openCount-1, extraLeft, extraRight, '#', path, res)
        }
    } else {
        path = append(path, s[i])
        removeParentheses(s, i+1, openCount, extraLeft, extraRight, '#', path, res)
    }
}

func getExtraParentheses(s string) (int, int) {
    var left, right int
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else if s[i] == ')'{
            if left > 0 {
                left--
            } else {
                right++
            }
        }
    }
    return left, right
}

---


func removeInvalidParentheses(s string) []string {
    left, right := computeExtraParentheses(s)
    var res []string
    removeParentheses(s, 0, 0, left, right, '*', nil, &res)
    return res
}

func removeParentheses(
    s string, i, open, extraLeft, extraRight int, lastSkipped byte, path []byte, res *[]string) {
    if extraLeft < 0 || extraRight < 0 || open < 0 {
        return
    } else if i == len(s) {
        if extraLeft == 0 && extraRight == 0 && open == 0 {
            *res = append(*res, string(path))
        }
        return
    }
    if s[i] == '(' {
        removeParentheses(s, i+1, open, extraLeft-1, extraRight, '(', path, res)
        if lastSkipped != '(' {
            path = append(path, s[i])
            removeParentheses(s, i+1, open+1, extraLeft, extraRight, '*', path, res)
        }
    } else if s[i] == ')' {
        removeParentheses(s, i+1, open, extraLeft, extraRight-1, ')', path, res)
        if lastSkipped != ')' {
            path = append(path, s[i])
            removeParentheses(s, i+1, open-1, extraLeft, extraRight, '*', path, res)
        }
    } else {
        path = append(path, s[i])
        removeParentheses(s, i+1, open, extraLeft, extraRight, '*', path, res)
    }
}

func computeExtraParentheses(s string) (int, int) {
    var left, right int
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else if s[i] == ')' {
            if left > 0 {
                left--
            } else {
                right++
            }
        }
    }
    return left, right
}

---

/*
    ()())()
    ( ())()
    ()( )()
 */
func removeInvalidParentheses(input string) []string {
    set := map[string]struct{}{}
    minRemoved := math.MaxInt32
    removeHelper(input, 0, 0, 0, 0, []byte{}, &minRemoved, &set)
    var res []string
    for expr := range set {
        res = append(res, expr)
    }
    return res
}

func removeHelper(s string, i, left, right, removed int, expr []byte,
                  minRemoved *int, set *map[string]struct{}) {
    if i == len(s) {
        if left == right && removed <= *minRemoved {
            if removed < *minRemoved {
                *set = map[string]struct{}{}
                *minRemoved = removed
            }
            (*set)[string(expr)] = struct{}{}
        }
        return
    }
    if s[i] != '(' && s[i] != ')' {
        removeHelper(s, i+1, left, right, removed,
                     append(expr, s[i]), minRemoved, set)
    } else {
        removeHelper(s, i+1, left, right, removed+1, expr, minRemoved, set)
        if s[i] == '(' {
            removeHelper(s, i+1, left+1, right, removed,
                         append(expr, s[i]), minRemoved, set)
        } else if right < left {
            removeHelper(s, i+1, left, right+1, removed,
                         append(expr, s[i]), minRemoved, set)
        }
    }
}

---

func removeInvalidParentheses(s string) []string {
    left, right := extraParentheses(s)
    var res []string
    dfs(s, 0, 0, left, right, '*', []byte{}, &res)
    return res
}

func dfs(s string, i, open, left, right int, skipped byte, tmp []byte, res *[]string) {
    if left < 0 || right < 0 || open < 0 { return }
    if i == len(s) {
        if left == 0 && right == 0 && open == 0 {
            *res = append(*res, string(tmp))
        }
        return
    }
    if s[i] == '(' {
        // Skip it
        dfs(s, i+1, open, left-1, right, '(', tmp, res)
        // Take it
        if skipped != '(' {
            dfs(s, i+1, open+1, left, right, '*', append(tmp, s[i]), res)
        }
    } else if s[i] == ')' {
        // Skip it
        dfs(s, i+1, open, left, right-1, ')', tmp, res)
        // Take it
        if open > 0 && skipped != ')' {
            dfs(s, i+1, open-1, left, right, '*', append(tmp, s[i]), res)
        }
    } else {
        dfs(s, i+1, open, left, right, '*', append(tmp, s[i]), res)
    }
}

func extraParentheses(s string) (int, int) {
    var left, right int
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else if s[i] == ')' {
            if left > 0 {
                left--
            } else {
                right++
            }
        }
    }
    return left, right
}

---

func removeInvalidParentheses(s string) []string {
    var res []string
    removeHelper(s, 0, 0, &res, "()")
    return res
}

func removeHelper(s string, lastI, lastJ int, res *[]string, order string) {
    stack := 0
    for i := lastI; i < len(s); i++ {
        if s[i] == order[0] { stack++ }
        if s[i] == order[1] { stack-- }
        if stack >= 0 { continue }
        for j := lastJ; j <= i; j++ {
            if s[j] == order[1] && (j == lastJ || s[j-1] != order[1]) {
                removeHelper(s[:j] + s[j+1:], i, j, res, order)
            }
        }
        return
    }
    if order[0] == '(' {
        removeHelper(reverse(s), 0, 0, res, ")(")
    } else {
        *res = append(*res, reverse(s))
    }
}

func reverse(s string) string {
    buffer := []byte(s)
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        buffer[i], buffer[j] = buffer[j], buffer[i]
    }
    return string(buffer)
}
