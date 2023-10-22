func romanToInt(s string) int {
    table := map[byte]int{
        'I': 1,     'V': 5,     'X': 10,
        'L': 50,    'C': 100,   'D': 500,   'M': 1000,
    }
    var result, curr, prev int
    for i := 0; i < len(s); i++ {
        curr = table[s[i]]
        result += curr
        if prev < curr {
            result -= 2 * prev
        }
        prev = curr
    }
    return result
}
