func groupStrings(strings []string) [][]string {
    groupMap := map[string][]string{}
    for _, s := range strings {
        key := getKey(s)
        groupMap[key] = append(groupMap[key], s)
    }
    var res [][]string
    for _, group := range groupMap {
        res = append(res, group)
    }
    return res
}

func getKey(s string) string {
    var res []byte
    offset := s[0] - 'a'
    for i := 0; i < len(s); i++ {
        ch := s[i] - offset
        if ch < 'a' { ch += 26 }
        res = append(res, ch)
    }
    return string(res)
}

---

func groupStrings(strings []string) [][]string {
    groups := map[string][]string{}
    for _, str := range strings {
        id := groupId(str)
        groups[id] = append(groups[id], str)
    }
    var res [][]string
    for _, groupedStrings := range groups {
        res = append(res, groupedStrings)
    }
    return res
}

func groupId(s string) string {
    var bs []byte
    var i int
    for ; i < len(s)-1; i++ {
        next := s[i+1] + 26
        bs = append(bs, (next-s[i])%26)
    }
    return string(bs)
}

---

func groupStrings(strings []string) [][]string {
    groups := map[string][]string{}
    for _, str := range strings {
        id := groupId(str)
        groups[id] = append(groups[id], str)
    }
    var res [][]string
    for _, groupedStrings := range groups {
        res = append(res, groupedStrings) 
    }
    return res
}

func groupId(s string) string {
    var bs []byte
    var i int
    for ; i < len(s)-1; i++ {
        next := s[i+1] + 26
        bs = append(bs, (next-s[i])%26) 
    }
    return string(bs)
}

func groupStrings(strings []string) [][]string {
    groupMap := make(map[string][]string)
    for i := 0; i < len(strings); i++ {
        s := key(strings[i])
        groupMap[s] = append(groupMap[s], strings[i])
    }
    result := make([][]string, len(groupMap))
    j := 0
    for _, list := range groupMap {
        result[j] = list
        j++
    }
    return result
}

func key(s string) string {
    const numChars int = 'z' - 'a' + 1
    keyStr := []byte{}
    for i := 1; i < len(s); i++ {
        diff := int(s[i]) - int(s[i - 1])
        if diff < 0 { diff += numChars }
        keyStr = append(keyStr, 'a' + byte(diff))
    }
    return string(keyStr)
}
