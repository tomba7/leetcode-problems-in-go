func countAndSay(n int) string {
    if n < 0 {return ""}
    result, tmp := []byte{'1'}, []byte{}
    for i := 2; i <= n; i++ {
        count := 0
        for j := 0; j < len(result); j++ {
            count++
            if j + 1 == len(result) || result[j] != result[j + 1] {
                tmp = append(tmp, fmt.Sprintf("%d%c", count, result[j])...)
                count = 0
            }
        }
        result, tmp = tmp, nil
    }
    return string(result)
}
