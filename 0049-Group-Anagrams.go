func groupAnagrams(strs []string) [][]string {
    groupMap := make(map[string][]string)
    for i := 0; i < len(strs); i++ {
        b := Bytes(strs[i])
        sort.Sort(b)
        groupMap[string(b)] = append(groupMap[string(b)], strs[i])
    }
    result := make([][]string, len(groupMap))
    i := 0
    for _, agGroup := range groupMap {
        result[i] = agGroup
        i++
    }
    return result
}

type Bytes []byte
func (b Bytes) Less(i, j int) bool { return b[i] < b[j] }
func (b Bytes) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b Bytes) Len() int { return len(b) }
