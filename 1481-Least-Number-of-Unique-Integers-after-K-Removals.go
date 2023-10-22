// With Buckets
func findLeastNumOfUniqueInts(arr []int, k int) int {
    freq := map[int]int{}
    for _, a := range arr {
        freq[a]++
    }
    buckets := make([][]int, len(arr)+1)
    for key, count := range freq {
        buckets[count] = append(buckets[count], key)
    }
    for count := 0; count <= len(arr); count++ {
        if k == 0 { break }
        for len(buckets[count]) != 0 && k >= count {
            n := len(buckets[count])
            top := buckets[count][n-1]
            buckets[count] = buckets[count][:n-1]
            delete(freq, top)
            k -= count
        }
    }
    return len(freq)
}

---
// Straight Forward Solution
// [5,5,5,5,5,2,2,2,3,3,1,6,6,7] k = 7
// (1, 1), (7, 1), (6, 2), (3, 2), (2, 3), (5, 5)
func findLeastNumOfUniqueInts(arr []int, k int) int {
    freq := make(map[int]int)
    for _, num := range arr {
        freq[num]++
    }
    var pairs [][]int
    for num, count := range freq {
        pairs = append(pairs, []int{num, count})
    }
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][1] < pairs[j][1]
    })
    for i := 0; i < len(pairs); i++ {
        if k > 0 && k >= pairs[i][1] {
            k -= pairs[i][1]
        } else {
            return len(pairs) - i
        }
    }
    return 0
}

---

// Complicated O(n) solution
func findLeastNumOfUniqueInts(arr []int, k int) int {
    count := map[int]int{}
    for _, a := range arr {
        count[a]++
    }
    remaining, occur := len(count), 1
    occurenceCount := make([]int, len(arr)+1)
    for _, val := range count {
        occurenceCount[val]++
    }
    for k > 0 {
        if k - occur*occurenceCount[occur] >= 0 {
            k -= occur * occurenceCount[occur]
            remaining -= occurenceCount[occur]
            occur++
        } else {
            return remaining - k/occur
        }
    }
    return remaining
}
