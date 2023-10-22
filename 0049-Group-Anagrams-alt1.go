func groupAnagrams(strs []string) (result [][]string) {
    group := map[string][]string{}
    for i := range strs {
        s := []byte(strs[i])
        sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
        group[string(s)] = append(group[string(s)], strs[i])
    }
    for _, list := range group { result = append(result, list) }
    return
}
