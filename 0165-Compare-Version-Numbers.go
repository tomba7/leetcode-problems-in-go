func compareVersion(v1, v2 string) int {
    var i, j int
    for i < len(v1) || j < len(v2) {
        num1, num2 := 0, 0
        for i < len(v1) && v1[i] != '.' {
            num1 = num1 * 10 + int(v1[i] - '0')
            i++
        }
        for j < len(v2) && v2[j] != '.' {
            num2 = num2 * 10 + int(v2[j] - '0')
            j++
        }
        if num1 > num2 {
            return 1
        } else if num1 < num2 {
            return -1
        }
        i++
        j++
    }
    return 0
}
