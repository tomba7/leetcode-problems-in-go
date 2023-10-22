func lengthLongestPath(input string) int {
    stack := []int{0}
    maxLen := 0
    for _, s := range strings.Split(input, "\n") {
        level := strings.Count(s, "\t")
        for level < len(stack) - 1 {
            stack = stack[:len(stack) - 1]
        }
        // Lets remove "\t" and add"/"
        length := stack[len(stack) - 1] + len(s) - level + 1;
        stack = append(stack, length)
        if strings.Contains(s, ".") {
            maxLen = max(maxLen, length - 1)
        }
    }
    return maxLen
}

func max(x, y int) int { if x > y { return x }; return y }
