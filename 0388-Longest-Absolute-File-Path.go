func lengthLongestPath(input string) int {
    stack := list.New()
    var maxLen, running int
    for _, s := range strings.Split(input, "\n") {
        level := strings.Count(s, "\t")
        for level != stack.Len() {
            running -= stack.Back().Value.(int)
            stack.Remove(stack.Back())
        }
        curLen := len(s) - level + 1
        stack.PushBack(curLen)
        running += curLen
        if strings.Contains(s, ".") {
            maxLen = max(maxLen, running - 1)
        }
    }
    return maxLen
}

func max(x, y int) int { if x > y { return x }; return y }
