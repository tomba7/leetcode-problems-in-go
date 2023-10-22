/*
    ")()())"
          ^
    l = 0
    r = 0
    m = 4
 */
func longestValidParentheses(s string) int {
    var left, right, res int
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            res = max(res, left + right)
        } else if left < right {
            left, right = 0, 0
        }
    }
    left, right = 0, 0
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            res = max(res, left + right)
        } else if left > right {
            left, right = 0, 0
        }
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }

---

/*
    )()())
         ^
    5
    res = 4-0 = 4

    ())((())
    ^
 */
func longestValidParentheses(s string) int {
    var res int
    stack := list.New()
    stack.PushBack(-1)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack.PushBack(i)
        } else {
            stack.Remove(stack.Back())
            if stack.Len() == 0 {
                stack.PushBack(i)
            } else {
                res = max(res, i - stack.Back().Value.(int))
            }
        }
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }
