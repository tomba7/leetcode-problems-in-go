// Time: O(n log n), Space: O(n)
func frequencySort(s string) string {
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
    }
    var pairs []pair
    for ch, count := range freq {
        pairs = append(pairs, pair{
            ch: ch, count: count,
        })
    }
    sort.Slice(pairs, func(i, j int) bool{
        return pairs[i].count > pairs[j].count
    })
    var res []byte
    for _, pair := range pairs {
        for i := 0; i < pair.count; i++ {
            res = append(res, pair.ch)
        }
    }
    return string(res)
}

type pair struct {
    ch byte
    count int
}

---

// Time: O(n), Space: O(n)
func frequencySort(s string) string {
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
    }
    var n int
    for _, count := range freq {
        if count > n { n = count }
    }
    bucket := make([][]byte, n+1)
    for ch, count := range freq {
        bucket[count] = append(bucket[count], ch)
    }
    var res []byte
    for i := n; i >= 0; i-- {
        for j := 0; j < len(bucket[i]); j++ {
            for k := 0; k < i; k++ {
                res = append(res, bucket[i][j])
            }
        }
    }
    return string(res)
}
