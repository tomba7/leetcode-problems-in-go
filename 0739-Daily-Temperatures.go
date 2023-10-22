func dailyTemperatures(temps []int) []int {
    n := len(temps)
    res := make([]int, n)
    s := list.New()
    for i := n-1; i >= 0; i-- {
        for s.Len() != 0 && s.Back().Value.([]int)[0] <= temps[i] {
            s.Remove(s.Back())
        }
        if s.Len() != 0 {
            res[i] = s.Back().Value.([]int)[1] - i
        }
        s.PushBack([]int{temps[i], i})
    }
    return res
}

---

/*
      0. 1. 2. 3. 4. 5. 6. 7
    [73,74,75,71,69,72,76,73]
            ^
            4   2  1  1  0  0                     
 S = | 6 |
 - Use a stack and iterate from the end of the array
 - Peek and pop the stack until the top of the stack is
   greater than the current element
 - If the stack is empty then answer[i] = 0
 - If not then answer[i] = topIndex - currIndex
*/
func dailyTemperatures(temps []int) []int {
    s := list.New()
    answers := make([]int, len(temps))
    for i := len(temps)-1; i >= 0; i-- {
        for s.Len() != 0 && temps[i] >= temps[s.Back().Value.(int)] {
            s.Remove(s.Back())
        }
        if s.Len() == 0 {
            answers[i] = 0
        } else {
            answers[i] = s.Back().Value.(int) - i 
        }
        s.PushBack(i)
    }
    return answers
}
