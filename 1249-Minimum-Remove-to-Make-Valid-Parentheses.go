func minRemoveToMakeValid(input string) string {
    s := list.New()
    n := len(input)
    removeSet := make(map[int]struct{})
    for i := 0; i < n; i++ {
        if input[i] == '(' {
            s.PushBack(i)
        } else if input[i] == ')'{
            if s.Len() == 0 {
                removeSet[i] = struct{}{}
            } else {
                s.Remove(s.Back())
            }
        }
    }
    for s.Len() != 0 {
        removeSet[s.Remove(s.Back()).(int)] = struct{}{}
    }
    var res []byte
    for i := 0; i < n; i++ {
        if _, exist := removeSet[i]; exist { continue }
        res = append(res, input[i])
    }
    return string(res)
}

---

func minRemoveToMakeValid(s string) string {
    res := reverse(
        removeInvalidBrackets(s, '(', ')'),
    )
    res = reverse(
        removeInvalidBrackets(string(res), ')', '('),
    )
    return string(res)
}

func removeInvalidBrackets(s string, opener, closer byte) []byte {
    var res []byte
    var balance int
    for i := 0; i < len(s); i++ {
        if s[i] == opener {
            balance++
        } else if s[i] == closer {
            if balance == 0 {
                continue
            }
            balance--
        }
        res = append(res, s[i])
    }
    return res
}

func reverse(s []byte) []byte {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}
