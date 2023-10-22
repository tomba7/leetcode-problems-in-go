/*
    l = 3, r = 3        0
                        
                     /      \
                  2,3         3,2   
                  (            )
                 
              1,3     2,2
              (        )

 - if left > right { return }
 */
func generateParenthesis(n int) []string {
    var res []string
    generateParenthesisHelper(n, n, []byte{}, &res)
    return res
}

func generateParenthesisHelper(left, right int, path []byte, res *[]string) {
    if left > right || left < 0 { return }
    if left == 0 && right == 0 {
        *res = append(*res, string(path))
        return
    }
    generateParenthesisHelper(left-1, right, append(path, '('), res)
    generateParenthesisHelper(left, right-1, append(path, ')'), res)
}

---

func generateParenthesis(n int) []string {
    var res []string
    helper(n, n, make([]byte, n*2), 0, &res)
    return res
}

func helper(left, right int, buf []byte, i int, res *[]string) {
    if left == 0 && right == 0 {
        *res = append(*res, string(buf))
        return
    } else if left > right || left < 0 {
        return
    }
    buf[i] = '('
    helper(left-1, right, buf, i+1, res)
    buf[i] = ')'
    helper(left, right-1, buf, i+1, res)
}

---

func generateParenthesis(n int) []string {
    result := []string{}
    generateParenthesisHelper(n, n, []byte{}, &result)
    return result
}

func generateParenthesisHelper(left, right int, combo []byte, result *[]string) {
    if left < 0 || left > right {
        return
    } else if left == 0 && right == 0 {
        *result = append(*result, string(combo))
        return
    }
    generateParenthesisHelper(left-1, right, append(combo, '('), result)
    generateParenthesisHelper(left, right-1, append(combo, ')'), result)
}
