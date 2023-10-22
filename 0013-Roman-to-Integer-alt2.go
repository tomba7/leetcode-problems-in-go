func romanToInt(s string) int {
    table := map[byte]int{
        'I': 1,     'V': 5,     'X': 10,
        'L': 50,    'C': 100,   'D': 500,   'M': 1000,
    }
    result := 0
    for i := 0; i < len(s); i++ {
        result += table[s[i]]
        if i != 0 && table[s[i-1]] < table[s[i]] {
            result -= 2 * table[s[i-1]]
        }
    }
    return result
}
