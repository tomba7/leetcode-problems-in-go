func fractionToDecimal(numerator, denominator int) string {
    if numerator == 0 { return "0" }
    var res []byte
    if numerator > 0 && denominator < 0 || numerator < 0 && denominator > 0 {
        res = append(res, '-')
    }
    numerator = abs(numerator)
    denominator = abs(denominator)
    res = append(res, strconv.Itoa(numerator/denominator)...)
    numerator %= denominator
    if numerator == 0 {
        return string(res)
    }
    res = append(res, '.')
    indexMap := map[int]int{}
    indexMap[numerator] = len(res)
    for numerator != 0 {
        numerator *= 10
        res = append(res, strconv.Itoa(numerator/denominator)...)
        numerator %= denominator
        if _, contains := indexMap[numerator]; contains {
            i := indexMap[numerator]
            res = append(res[:i+1], res[i:]...)
            res[i] = '('
            res = append(res, ')')
            break
        } else {
            indexMap[numerator] = len(res)
        }
    }
    return string(res)
}

func abs(x int) int {
    if x < 0 { return -x }; return x
}
