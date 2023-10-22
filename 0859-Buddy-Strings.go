/*
    Algo
    - If len(s) != len(goal) return false
    - If s == goal, then there needs to be duplicate chars
    - If all other cases
      * s[i] != goal[i]
      * And the num er of indices where s[i] != goal[i] < 2
      * And s[i] == goal[j] && s[j] == goal[i]
 */
func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) { return false } 
    n := len(s)
    if s == goal {
        set := map[byte]struct{}{}
        for i := 0; i < n; i++ {
            set[s[i]] = struct{}{}
        }
        return len(set) < n
    } 
    var diff []int
    for i := 0; i < n; i++ {
        if s[i] != goal[i] {
            diff = append(diff, i)
        }
    }
    return len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
}
