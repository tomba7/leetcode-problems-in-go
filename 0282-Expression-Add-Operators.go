/*
    123     t = 6
    1+2+3 
    1*2*3
    
    105     t = 5
    10-5
    1*0+5
    
    222     t = 24
    22+2
    2+22
    
    2345    t = 19
    2+3*4+5
   
    0           1           2           3
    2.    +     3.    *     4.    +     5
   (2, 2)      (5, 3)      (14,12)     (19, 5) 
                        5 - 3 + 4 * 3
    
    23452   t = 12
    2*3-4+5*2
    
    0           1           2           3           4
    2.    *     3.    -     4.    +     5     *     2
   (2, 2)      (6, 6)      (2, -4)     (7, 5)
          2 - 2 + 2 * 3                         7 - 5 + 5 * 2 
    
    1234
    i = 0, [1, 12, 123, 1234]
               (1)              (12)            (123)
    i = 1, [2, 23, 234]        [3, 34]           [4]
 */
func addOperators(num string, target int) []string {
    var res []string
    helper(num, target, 0, "", 0, 0, &res)
    return res
}

func helper(num string, target, i int, expr string, val, prev int, res *[]string) {
    if i == len(num) {
        if target == val { *res = append(*res, expr) }
        return
    }
    for j := i; j < len(num); j++ {
        currStr := num[i:j+1]
        curr, _ := strconv.Atoi(currStr)
        if i != j && num[i] == '0' { break }
        if i == 0 {
            helper(num, target, j+1, currStr, curr, curr, res)
        } else {
            helper(num, target, j+1, expr+"+"+currStr, val+curr, curr, res)
            helper(num, target, j+1, expr+"-"+currStr, val-curr, -curr, res)
            helper(num, target, j+1, expr+"*"+currStr, val-prev+prev*curr, prev*curr, res)
        }
    }
}

---

func addOperators(s string, target int) []string {
    var res []string
    helper(s, 0, "", 0, 0, target, &res)
    return res
}

func helper(s string, i int, expr string, total, prev, target int, res *[]string) {
    if i == len(s) {
        if total == target { *res = append(*res, expr) }
        return
    }
    for j := i; j < len(s); j++ {
        numStr := s[i:j+1]
        num, _ := strconv.Atoi(numStr)
        if len(numStr) > 1 && numStr[0] == '0' { continue }
        if i == 0 {
            helper(s, j+1, numStr, num, num, target, res)
        } else {
            helper(s, j+1, expr+"*"+numStr, total-prev+prev*num, prev*num, target, res)
            helper(s, j+1, expr+"+"+numStr, total+num, num, target, res)
            helper(s, j+1, expr+"-"+numStr, total-num, -num, target, res)
        }
    }
}

---

func addOperators(num string, target int) []string {
    var res []string
    helper(num, target, 0, 0, 0, 0, "", &res)
    return res
}

func helper(nums string, target, indx, prev, curr, val int, expr string, res *[]string) {
    if indx == len(nums) {
        if val == target && curr == 0 {
            *res = append(*res, expr[1:])
        }
        return
    }
    curr = curr * 10 + int(nums[indx]) - '0'
    valRep := strconv.Itoa(curr)
    if curr > 0 {
        helper(nums, target, indx+1, prev, curr, val, expr, res)
    }
    helper(nums, target, indx+1, curr, 0, val+curr, expr+"+"+valRep, res)
    if len(expr) > 0 {
        helper(nums, target, indx+1, -curr, 0, val-curr, expr+"-"+valRep, res)
        helper(nums, target, indx+1, curr*prev, 0, val-prev+(curr*prev), expr+"*"+valRep, res)
    }
}

---

func addOperators(nums string, target int) []string {
    var res []string
    helper(nums, 0, 0, 0, 0, target, []string{}, &res)
    return res
}

func helper(nums string, i, prev, curr, val, target int, ops []string, res *[]string) {
    if i == len(nums) {
        if val == target && curr == 0 { *res = append(*res, format(ops)) }
        return
    }
    curr = curr * 10 + int(nums[i] - '0')
    currStr := strconv.Itoa(curr)
    if curr > 0 {
        helper(nums, i+1, prev, curr, val, target, ops, res)
    }
    helper(nums, i+1, curr, 0, val+curr, target, append(ops, "+", currStr), res)
    if len(ops) > 0 {
        helper(nums, i+1, -curr, 0, val-curr, target, append(ops, "-", currStr), res)
        helper(nums, i+1, curr*prev, 0, val-prev+curr*prev, target, append(ops, "*", currStr), res)
    }
}

func format(ops []string) string {
    var res []byte
    for i := 1; i < len(ops); i++ {
        res = append(res, ops[i]...)
    }
    return string(res)
}
