/*
 - create a graph of [digit] -> [possible letters]
 - iterate the string and for every digit
 - loop through all possible letters
 - then recurse on the next digit
 - At the base, collect all possible combinations
*/
var graph = [][]byte{
    {},{},
    {'a', 'b', 'c'},        // 2
    {'d', 'e', 'f'},
    {'g', 'h', 'i'},
    {'j', 'k', 'l'},
    {'m', 'n', 'o'},
    {'p', 'q', 'r', 's'},
    {'t', 'u', 'v'},
    {'w', 'x', 'y', 'z'},   // 9
}

func letterCombinations(digits string) []string {
    var res []string
    letters := make([]byte, len(digits))
    letterCombinationsHelper(digits, 0, letters, &res)
    return res
}

func letterCombinationsHelper(digits string, i int, letters []byte, res *[]string) {
    if i == len(digits) {
        // You can avoid this check by checking for empty input strings in main
        // e.g. if len(digits) == 0 { return nil }
        if len(letters) != 0 {
            *res = append(*res, string(letters))
            return
        }
        return
    }
    digit := digits[i]
    for _, ch := range graph[digit-'0'] {
        letters[i] = ch
        letterCombinationsHelper(digits, i+1, letters, res)
    }
}

---

func letterCombinations(digits string) (result []string) {
    if len(digits) == 0 { return result }
    table := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
    letterCombinationsHelper(digits, table, 0, []byte{}, &result)
    return result
}

func letterCombinationsHelper(digits string, table []string, index int, combination []byte, result *[]string) {
    if index == len(digits) {
        *result = append(*result, string(combination))
        return
    }
    word := table[digits[index] - '0']
    for i := 0; i < len(word); i++ {
        letterCombinationsHelper(digits, table, index + 1, append(combination, word[i]), result)
    }
}
