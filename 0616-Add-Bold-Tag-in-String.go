/*
    Time: O(N*M*W), where N = len(nums), M = len(words), W = len(words[0])

    "a a a b b c c"
    "T T T T T T F"
     ^
 */
func addBoldTag(s string, words []string) string {
    boldChar := make([]bool, len(s))
    var end int
    for i := 0; i < len(s); i++ {
        for _, word := range words {
            if strings.HasPrefix(s[i:], word) {
                end = max(end, i + len(word))
            }
        }
        boldChar[i] = i < end
    }
    fmt.Println(boldChar)
    var res []byte
    for i := 0; i < len(s); i++ {
        if !boldChar[i] {
            res = append(res, s[i]) 
            continue
        }
        if i == 0 || !boldChar[i-1] {
            res = append(res, "<b>"...)
        }
        res = append(res, s[i])
        if i + 1 == len(s) || !boldChar[i+1] {
            res = append(res, "</b>"...)
        }
    }
    return string(res)
}

func max(x, y int) int { if x > y { return x }; return y }

---

// Brute force
func addBoldTag(s string, words []string) string {
    set := make(map[string]bool)
    for _, word := range words {
        set[word] = true
    }
    var ranges [][]int
    findWordRanges(s, 0, set, &ranges)

    stack := list.New()
    if len(ranges) != 0 {
        sort.Slice(ranges, func(i, j int) bool {
            return ranges[i][0] < ranges[j][0]
        })
        stack.PushBack(ranges[0])
        for i := 1; i < len(ranges); i++ {
            top := stack.Back().Value.([]int)
            s1, e1 := top[0], top[1]
            s2, e2 := ranges[i][0], ranges[i][1]
            if s2 <= e1 + 1 {
                stack.Remove(stack.Back())
                stack.PushBack([]int{s1, max(e1, e2)})
            } else {
                stack.PushBack(ranges[i])
            }
        }
    }
    var res string
    prev := 0
    for stack.Len() != 0 {
        curr := stack.Remove(stack.Front()).([]int)
        start, end := curr[0], curr[1]
        res = res + s[prev:start] + "<b>" + s[start:end+1] + "</b>"
        prev = end + 1
    }
    if prev != len(s) {
        res = res + s[prev:]
    }
    return res
}

func findWordRanges(s string, i int, set map[string]bool, ranges *[][]int) {
    if i == len(s) { return }
    for j := i; j < len(s); j++ {
        if set[s[i:j+1]] {
            *ranges = append(*ranges, []int{i, j})
        }
        if !seen[s[j+1:]] {
            findWordRanges(s, j+1, set, ranges)
        }
    }
}

func max(x, y int) int { if x > y { return x }; return y }
