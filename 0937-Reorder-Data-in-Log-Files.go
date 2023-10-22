/*
  Time Complexity:  O(M⋅NlogN)
  Space Complexity: O(M⋅logN)
 */
---
func reorderLogFiles(logs []string) []string {
    sort.SliceStable(logs, func(i, j int) bool {
        s1, s2 := logs[i], logs[j]
        x, y := strings.Index(s1, " "), strings.Index(s2, " ")
        log1, log2 := s1[x+1:], s2[y+1:]
        digit1, digit2 := digit(log1[0]), digit(log2[0])
        if digit1 && digit2 {
            return false
        } else if digit1 && !digit2 {
            return false
        } else if !digit1 && digit2 {
            return true
        } else {
            if log1 == log2 { return s1[:x] < s2[:y] }
            return log1 < log2
        }
    })
    return logs
}

func digit(ch byte) bool { return '0' <= ch && ch <= '9' }
