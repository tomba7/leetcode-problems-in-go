func characterReplacement(s string, k int) int {
    freq := map[byte]int{}
    var res, left, mostFreqLetter int
    for right := 0; right < len(s); right++ {
        freq[s[right]]++
        mostFreqLetter = max(mostFreqLetter, freq[s[right]])
        lettersToChange := right - left + 1 - mostFreqLetter
        if lettersToChange > k {
            freq[s[left]]--
            left++
        }
        res = max(res, right - left + 1)
    }
    return res
}

func max(x, y int) int { if x > y { return x }; return y }
