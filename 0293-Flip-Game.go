func generatePossibleNextMoves(s string) []string {
    result := []string{}
    for i := 1; i < len(s); i++ {
        if s[i] == '+' && s[i] == s[i - 1] {
            result = append(result, s[:i - 1] + "--" + s[i + 1:])
        }
    }
    return result
}
