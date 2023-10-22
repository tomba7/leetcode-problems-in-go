func addBinary(a string, b string) string {
    if len(a) < len(b) { return addBinary(b, a) }
    i, j := len(a)-1, len(b)-1
    var carry byte
    var res []byte
    for i >= 0 || j >= 0 || carry != 0 {
        if i >= 0 {
            carry += a[i] - '0'
            i--
        }
        if j >= 0 {
            carry += b[j] - '0'
            j--
        }
        res = append(res, '0' + carry%2)
        carry /= 2
    }
    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
    return string(res)
}

---

func addBinary(a, b string) string {
    result := []byte{}
    sum := 0
    for i, j := len(a) - 1, len(b) - 1; i >= 0 || j >= 0 || sum != 0; i, j = i - 1, j - 1 {
        if i >= 0 {sum += int(a[i] - '0')}
        if j >= 0 {sum += int(b[j] - '0')}
        result = append(result, byte(sum) % 2 + '0')
        sum /= 2
    }
    for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
        result[i], result[j] = result[j], result[i]
    }
    return string(result)
}

---

func addBinary(a, b string) string {
    var result []byte
    var carry byte
    for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
        if i >= 0 { carry += a[i]-'0' }
        if j >= 0 { carry += b[j]-'0' }
        result = append(result, carry%2 + '0')
        carry /= 2
    }
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    return string(result)
}

---

func addBinary(a string, b string) string {
    var sum []byte
    var c byte
    i, j := len(a)-1, len(b)-1
    for ; i >= 0 || j >= 0 || c != 0; i, j = i-1, j-1 {
        if i >= 0 { c += a[i]-'0' }
        if j >= 0 { c += b[j]-'0' }
        sum = append(sum, '0' + c%2)
        c /= 2
    }
    for i, j = 0, len(sum)-1; i < j; i, j = i+1, j-1 {
        sum[i], sum[j] = sum[j], sum[i]
    }
    return string(sum)
}
